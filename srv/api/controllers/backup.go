package apicontroller

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ranjbar-dev/backup-api/internal/config"
)

type StoreBackupRequest struct {
	File    *multipart.FileHeader `form:"file"`
	Project string                `form:"project"`
}

func (controller *Controller) StoreBackup(c *gin.Context) {

	var req StoreBackupRequest
	if err := c.ShouldBind(&req); err != nil {
		controller.error(c, err)
		return
	}

	if req.File == nil {
		controller.badRequest(c, "invalid file")
		return
	}

	if req.Project == "" {
		controller.badRequest(c, "invalid project name")
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		controller.error(c, err)
		return
	}

	// Check for directory existence and create it if needed
	backupDir := pwd + config.Single.BackupLocation()
	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		err = os.MkdirAll(backupDir, 0755) // Create directory with appropriate permissions
		if err != nil {
			controller.error(c, fmt.Errorf("failed to create backup directory: %w", err))
			return
		}
	}

	// Generate a unique filename with extension
	filename := filepath.Join(backupDir, fmt.Sprintf("%s-%d-%s", req.Project, time.Now().Unix(), req.File.Filename))

	// Open the uploaded file
	file, err := req.File.Open()
	if err != nil {
		controller.error(c, err)
		return
	}
	defer file.Close()

	// Create a new destination file
	dst, err := os.Create(filename)
	if err != nil {
		controller.error(c, fmt.Errorf("failed to create destination file: %w", err))
		return
	}
	defer dst.Close()

	// Copy uploaded file content
	if _, err := io.Copy(dst, file); err != nil {
		controller.error(c, fmt.Errorf("failed to copy file content: %w", err))
		return
	}

	controller.ok(c, map[string]any{
		"data": 200,
	})
}
