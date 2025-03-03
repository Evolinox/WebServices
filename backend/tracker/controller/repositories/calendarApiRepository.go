package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"tracker/infrastructure/dto/model"
)

type CalendarAPIRepository struct {
	baseURL string
}

func NewCalendarAPIRepository(baseURL string) *CalendarAPIRepository {
	return &CalendarAPIRepository{baseURL: baseURL}
}

func (r *CalendarAPIRepository) GetCalendarByDate(date string) ([]model.CalendarDTO, error) {
	url := fmt.Sprintf("%s/calendar/%s", r.baseURL, date)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch calendar entries: %s", resp.Status)
	}

	var calendars []model.CalendarDTO
	if err := json.NewDecoder(resp.Body).Decode(&calendars); err != nil {
		return nil, err
	}
	return calendars, nil
}

func (r *CalendarAPIRepository) AddCalendarEntry(entry model.CalendarDTO) error {
	url := fmt.Sprintf("%s/calendar/", r.baseURL)
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to add calendar entry: %s", resp.Status)
	}
	return nil
}

func (r *CalendarAPIRepository) UpdateCalendarEntry(id string, entry model.CalendarDTO) error {
	url := fmt.Sprintf("%s/calendar/%s", r.baseURL, id)
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update calendar entry: %s", resp.Status)
	}
	return nil
}

func (r *CalendarAPIRepository) DeleteCalendar(id string) error {
	url := fmt.Sprintf("%s/calendar/%s", r.baseURL, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("calendar entry not found")
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete calendar entry: %s", resp.Status)
	}
	return nil
}
