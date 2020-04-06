# Тестирование сервиса calendar
# Тестируем операции:
#  1) Создание события (Event);
#  2) Изменение события;
#  3) Удаление события;
#
Feature: service calendar
	In order to test the calendar application
	As an GRPC client operates with the service through API
	The service should be able to do the following

	Scenario: should add event
        When I call grpc calendar method AddEvent
		Then The error should be nil
		And The add response success should be true

	Scenario: Update event
		Given I have the Event
		When I call grpc calendar method UpdateEvent
		Then The error should be nil
		And The update response success should be true

	Scenario: Delete event
		Given I have the event ID
		When I call grpc calendar method DeleteEvent
		Then The error should be nil
		And The delete response success should be true
