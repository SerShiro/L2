package dev

import (
	"fmt"
	"log"
	"net/http"
)

func calendarAPI() {
	http.Handle("/create_event", loggingMiddleware(http.HandlerFunc(createEventHandler)))
	http.Handle("/update_event", loggingMiddleware(http.HandlerFunc(updateEventHandler)))
	http.Handle("/delete_event", loggingMiddleware(http.HandlerFunc(deleteEventHandler)))
	http.Handle("/events_for_day", loggingMiddleware(http.HandlerFunc(eventsForDayHandler)))
	http.Handle("/events_for_week", loggingMiddleware(http.HandlerFunc(eventsForWeekHandler)))
	http.Handle("/events_for_month", loggingMiddleware(http.HandlerFunc(eventsForMonthHandler)))

	port := ":8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
