package repository

import (
	"time"

	"github.com/khasyah-fr/phinco-attendance/models"
	"gorm.io/gorm"
)

type AttendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{
		db: db,
	}
}

func (ar *AttendanceRepository) GetAllAttendancesByUserID(userId int64, filter string) ([]models.Attendance, error) {
	var attendances []models.Attendance

	query := ar.db.Where("user_id = ?", userId)
	switch filter {
	case "day":
		query = query.Where("attend_time >= CURDATE()")
	case "week":
		query = query.Where("attend_time >= CURDATE() - INTERVAL DAYOFWEEK(CURDATE()) - 1 DAY")
	case "month":
		query = query.Where("attend_time >= DATE_FORMAT(NOW() ,'%Y-%m-01')")
	case "year":
		query = query.Where("attend_time >= DATE_FORMAT(NOW() ,'%Y-01-01')")
	}

	result := query.Find(&attendances)
	if result.Error != nil {
		return nil, result.Error
	}

	return attendances, nil
}

func (ar *AttendanceRepository) CheckIn(userId int64, locationId int64) error {
	attendance := models.Attendance{
		UserId:     userId,
		LocationId: locationId,
		Type:       "checkin",
		AttendTime: time.Now(),
	}

	result := ar.db.Create(&attendance)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ar *AttendanceRepository) CheckOut(userId int64, locationId int64) error {
	attendance := models.Attendance{
		UserId:     userId,
		LocationId: locationId,
		Type:       "checkout",
		AttendTime: time.Now(),
	}

	result := ar.db.Create(&attendance)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
