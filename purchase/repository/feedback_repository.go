package repository

import (
	. "MPPL-Modul-4-master/models/purchase"
	"MPPL-Modul-4-master/purchase"
	"github.com/jinzhu/gorm"
)

type feedbackRepository struct {
	Conn *gorm.DB
}

func NewFeedbackRepository(Conn *gorm.DB) purchase.RepositoryFeedback  {
	return &feedbackRepository{Conn}
}

func (pr *feedbackRepository) Fetch() (res []*Feedback, err error) {
	var feedback []*Feedback
	err = pr.Conn.Find(&feedback).Error

	if err != nil {
		return nil, err
	}

	return feedback, nil
}

func (pr *feedbackRepository) GetById(id uint) (*Feedback, error) {
	var feedback_ Feedback
	err := pr.Conn.Find(&feedback_, id).Error

	if err != nil {
		return nil, err
	}

	return &feedback_, nil
}

func (pr *feedbackRepository) Update(p *Feedback) error {
	var feedback_ Feedback
	pr.Conn.Find(feedback_, )

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *feedbackRepository) Store(p *Feedback) error {
	err := pr.Conn.Create(&p).Error
	if err != nil {
		return err
	}

	return nil
}

func (pr *feedbackRepository) Delete(id uint) error {
	var feedback_ Feedback
	pr.Conn.Find(&feedback_)
	err := pr.Conn.Delete(&feedback_).Error

	if err != nil {
		return err
	}

	return nil
}