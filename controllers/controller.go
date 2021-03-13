package controllers

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadBreackPoint(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		etx := filepath.Ext(file.Filename)
		filename := uuid.New().String() + etx
		if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "Error By Saving The File",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Uploaded!",
	})
}

func CreateDirBreakPoint(c *gin.Context) {
	path := c.Param("path")
	rpath := convertPath(path)
	err := os.Mkdir(rpath, os.ModePerm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "Error By Creating The Dir",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Created!",
	})
}

func ReadDirBreakPoint(c *gin.Context) {
	path := c.Param("path")
	rpath := convertPath(path)
	info, err := ioutil.ReadDir(rpath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "Error By Reading The Dir",
		})
		return
	}
	data := make([]string, len(info))

	for index, i := range info {
		data[index] = i.Name()
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func convertPath(oldPath string) string {
	num := strings.Count(oldPath, "-")
	newPath := strings.Replace(oldPath, "-", "/", num) + "/"
	rPath := path.Join("uploads/" + newPath)
	return rPath
}
