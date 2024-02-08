package dev

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Event struct {
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
}

var events []Event

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUserID(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid user_id")
		return
	}

	date, err := parseDate(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid date")
		return
	}

	for _, existingEvent := range events {
		if existingEvent.UserID == userID && existingEvent.Date == date {
			handleError(w, http.StatusConflict, "Event for this user and date already exists")
			return
		}
	}

	newEvent := Event{UserID: userID, Date: date}
	events = append(events, newEvent)

	responseJSON(w, http.StatusOK, map[string]interface{}{"result": "Event created successfully", "event": newEvent})
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUserID(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid user_id")
		return
	}

	oldDate, err := parseDate(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid date")
		return
	}

	newDate, err := parseNewDate(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid date")
		return
	}

	isExisted := false
	var indexNewEvent int
	for index, existingEvent := range events {
		if existingEvent.UserID == userID && existingEvent.Date == oldDate {
			isExisted = true
			indexNewEvent = index
			break
		}
	}

	if !isExisted {
		handleError(w, http.StatusServiceUnavailable, "Event for this user and date doesn't  exists")
		return
	}
	events[indexNewEvent].Date = newDate

	updatedEvent := Event{UserID: userID, Date: newDate}
	responseJSON(w, http.StatusOK, map[string]interface{}{"result": "Event updated successfully", "event": updatedEvent})
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUserID(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid user_id")
		return
	}

	date, err := parseDate(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid date")
		return
	}

	isExisted := false
	foundIndex := -1
	for i, existingEvent := range events {
		if existingEvent.UserID == userID && existingEvent.Date == date {
			isExisted = true
			foundIndex = i
			break
		}
	}

	if !isExisted {
		handleError(w, http.StatusServiceUnavailable, "Event for this user and date doesn't  exists")
		return
	}

	deletedEvent := events[foundIndex]
	events = append(events[:foundIndex], events[foundIndex+1:]...)

	responseJSON(w, http.StatusOK, map[string]interface{}{"result": "Event deleted successfully", "event": deletedEvent})
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUserID(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid user_id")
		return
	}

	date, err := parseDate(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid date")
		return
	}

	var eventsForDay []Event
	for _, existingEvent := range events {
		if existingEvent.UserID == userID && existingEvent.Date == date {
			eventsForDay = append(eventsForDay, existingEvent)
		}
	}

	responseJSON(w, http.StatusOK, map[string]interface{}{"result": "Events retrieved successfully", "events": eventsForDay})
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUserID(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid user_id")
		return
	}

	date, err := parseDate(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid date")
		return
	}

	startDate, endDate := calculateWeekDates(date)

	var eventsForWeek []Event
	for _, existingEvent := range events {
		if existingEvent.UserID == userID && existingEvent.Date >= startDate && existingEvent.Date <= endDate {
			eventsForWeek = append(eventsForWeek, existingEvent)
		}
	}

	responseJSON(w, http.StatusOK, map[string]interface{}{"result": "Events retrieved successfully", "events": eventsForWeek})
}
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := parseUserID(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid user_id")
		return
	}

	date, err := parseDate(r)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Invalid date")
		return
	}

	var eventsForMonth []Event
	currentEventDate, _ := time.Parse("2006-01-02", date)
	for _, existingEvent := range events {
		existingEventDate, _ := time.Parse("2006-01-02", existingEvent.Date)
		if existingEvent.UserID == userID && currentEventDate.Month() == existingEventDate.Month() {
			eventsForMonth = append(eventsForMonth, existingEvent)
		}
	}

	responseJSON(w, http.StatusOK, map[string]interface{}{"result": "Events retrieved successfully", "events": eventsForMonth})
}

func parseUserID(r *http.Request) (int, error) {
	userIDStr := r.FormValue("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func parseDate(r *http.Request) (string, error) {
	return r.FormValue("date"), nil
}
func parseNewDate(r *http.Request) (string, error) {
	return r.FormValue("new_date"), nil
}

func responseJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func handleError(w http.ResponseWriter, statusCode int, message string) {
	responseJSON(w, statusCode, map[string]interface{}{"error": message})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func calculateWeekDates(date string) (string, string) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", ""
	}

	weekday := parsedDate.Weekday()

	startDate := parsedDate.AddDate(0, 0, -int(weekday))
	endDate := parsedDate.AddDate(0, 0, 6-int(weekday))

	startDateStr := startDate.Format("2006-01-02")
	endDateStr := endDate.Format("2006-01-02")

	return startDateStr, endDateStr
}
