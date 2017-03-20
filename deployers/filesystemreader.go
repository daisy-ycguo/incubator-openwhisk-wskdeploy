/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package deployers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"io/ioutil"
	"archive/zip"

	"github.com/openwhisk/openwhisk-client-go/whisk"
	"github.com/openwhisk/openwhisk-wskdeploy/parsers"
	"github.com/openwhisk/openwhisk-wskdeploy/utils"
)

// name of directory that can contain source code
const FileSystemSourceDirectoryName = "actions"

type FileSystemReader struct {
	serviceDeployer *ServiceDeployer
	tempDir		string
}

func NewFileSystemReader(serviceDeployer *ServiceDeployer) *FileSystemReader {
	var reader FileSystemReader
	reader.serviceDeployer = serviceDeployer
	
	return &reader
}

func (reader *FileSystemReader) ReadProjectDirectory(manifest *parsers.ManifestYAML) ([]utils.ActionRecord, error) {

	fmt.Println("Inspecting project directory for actions....")

	projectPathCount, err := reader.getFilePathCount(reader.serviceDeployer.ProjectPath)
	utils.Check(err)

	actions := make([]utils.ActionRecord, 0)

	err = filepath.Walk(reader.serviceDeployer.ProjectPath, func(fpath string, f os.FileInfo, err error) error {
		if fpath != reader.serviceDeployer.ProjectPath {
			pathCount, err := reader.getFilePathCount(fpath)
			utils.Check(err)

			if !f.IsDir() {
				if pathCount-projectPathCount == 1 || strings.HasPrefix(fpath, reader.serviceDeployer.ProjectPath+"/"+FileSystemSourceDirectoryName) {
					ext := filepath.Ext(fpath)

					foundFile := false
					switch ext {
					case ".swift":
						foundFile = true
					case ".js":
						foundFile = true
					case ".py":
						foundFile = true
					}

					if foundFile == true {
						_, action, err := reader.CreateActionFromFile(reader.serviceDeployer.ManifestPath, fpath)
						utils.Check(err)

						var record utils.ActionRecord
						record.Action = action
						record.Packagename = manifest.Package.Packagename
						record.Filepath = fpath

						actions = append(actions, record)
					}
				}
			} else if strings.HasPrefix(fpath, reader.serviceDeployer.ProjectPath+"/"+FileSystemSourceDirectoryName) {
				fmt.Println("Zipping directory " + filepath.Base(fpath) + " for action source code.")
				_, action, err := reader.CreateActionFromDir(reader.serviceDeployer.ManifestPath, fpath)
			} else {
				return filepath.SkipDir
			}

		}
		return err
	})

	if err != nil {
		return nil, err
	}

	return actions, nil

}


func (reader *FileSystemReader) CreateActionFromFile(manipath, filePath string) (string, *whisk.Action, error) {
	ext := filepath.Ext(filePath)
	baseName := filepath.Base(filePath)
	name := strings.TrimSuffix(baseName, filepath.Ext(baseName))
	action := new(whisk.Action)

	// process source code files
	if ext == ".swift" || ext == ".js" || ext == ".py" {

		kind := "nodejs:default"

		switch ext {
		case ".swift":
			kind = "swift:default"
		case ".js":
			kind = "nodejs:default"
		case ".py":
			kind = "python"
		}

		dat, err := new(utils.ContentReader).LocalReader.ReadLocal(filePath)
		utils.Check(err)

		action.Exec = new(whisk.Exec)
		code := string(dat)
		action.Exec.Code = &code
		action.Exec.Kind = kind
		action.Name = name
		pub := false
		action.Publish = &pub
		return name, action, nil
	}
	// If the action is not supported, we better to return an error.
	return "", nil, errors.New("Unsupported action type.")
}

/**
This function will create an action with multiple files in a specific director.
The director will be compressed as a zip file.
The first meanful extension of those files under the directory will be used to decide the action type.
**/
func (reader *FileSystemReader) CreateActionFromDir(manipath, dir string) (string, *whisk.Action, error) {
	//open the dir
	baseName := filepath.Base(dir)
	dirfile, err := os.Open(dir)
	if err != nil {return "", nil, err}
	dirfileinfo, err := dirfile.Stat()
	if !dirfileinfo.IsDir() {return "", nil, errors.New(dir+" should be a directory.")}
	defer dirfile.Close()
	// get the first file extension as the action type
	kind := ""
	list, err := ioutil.ReadDir(dir)
found:
	for _, info := range list {
		if !info.IsDir() {
			filename := info.Name()
			ext := filepath.Ext(filename)
			switch ext {
			case ".swift":
				kind = "swift:default"
				break found
			case ".js":
				kind = "nodejs:default"
				break found
			case ".py":
				kind = "python"
				break found
			}
		}
	}
	if len(kind)==0 { return "", nil, errors.New("Unsupported action type.") }

	//create temp file
	tempfile := filepath.Join(getTempDir(), baseName+".zip")
	_, e := os.Stat(tempfile)
	if e!=nil { os.Remove(tempfile) }
	f,err := os.Create(tempfile)
	//compress the directory to a temp file
	err = compressDir(tempfile, dir)
	if err!=nil { return "", nil, err }

	//get the action type
	action := new(whisk.Action)
	dat, err := new(utils.ContentReader).LocalReader.ReadLocal(tempfile)
	utils.Check(err)
	action.Exec = new(whisk.Exec)
	code := string(dat)
	action.Exec.Code = &code
	action.Exec.Kind = kind
	action.Name = name
	pub := false
	action.Publish = &pub
	return name, action, nil	// process source code files
}

func (reader *FileSystemReader) getTempDir() string {
	if tempDir==nil {
		tempDir = ioutil.TempDir("","wskdeploy_temp")
	}
	return tempDir
}

func (reader *FileSystemReader) compressDir(destfile string, dirpath string) error {
	zipFile, _ := os.Create(destfile)
	defer zipFile.Close()
	myzip := zip.NewWriter(zipFile)
	defer myzip.Close()

	visit:=func(path string, f os.FileInfo, err error) error {
                if ( f == nil ) {return err}
                if f.IsDir() {
			return nil
		}

                thisfile, _ := os.Open(path)
		defer thisfile.Close()
		header, err := zip.FileInfoHeader(f)
		header.Name, _ = filepath.Rel(filepath.Dir(root), path)
		writer, err := myzip.CreateHeader(header)
		_, err = io.Copy(writer, thisfile)
		return err
        }

        err := filepath.Walk(dirpath, visit)
        if err != nil {
                fmt.Printf("filepath.Walk() returned %v\n", err)
		return err
        }
	myzip.Flush()
	return nil
}

func (reader *FileSystemReader) getFilePathCount(path string) (int, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return 0, err
	}

	pathList := strings.Split(absPath, "/")
	return len(pathList) - 1, nil
}

func (reader *FileSystemReader) SetFileActions(actions []utils.ActionRecord) error {

	dep := reader.serviceDeployer

	dep.mt.Lock()
	defer dep.mt.Unlock()

	for _, fileAction := range actions {
		existAction, exists := reader.serviceDeployer.Deployment.Packages[fileAction.Packagename].Actions[fileAction.Action.Name]

		if exists == true {
			if existAction.Filepath == fileAction.Filepath || existAction.Filepath == "" {
				// we're adding a filesystem detected action so just updated code and filepath if needed
				existAction.Action.Exec.Code = fileAction.Action.Exec.Code
				existAction.Filepath = fileAction.Filepath
			} else {
				// Action exists, but references two different sources
				return errors.New("Conflict detected for action named " + existAction.Action.Name + ". Found two locations for source file: " + existAction.Filepath + " and " + fileAction.Filepath)
			}
		} else {
			// not a new action so to actions in package
			reader.serviceDeployer.Deployment.Packages[fileAction.Packagename].Actions[fileAction.Action.Name] = fileAction
		}
	}

	return nil

}
