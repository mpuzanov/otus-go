INSERT INTO public.events("header", "text", "start_time", "end_time", "user_id", "reminder_before") VALUES 
('event 1', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '1 hour' ),
('event 2', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '0' ),
('event 3', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '0' ),
('event 4', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '30 minutes' ),
('event 5', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '30 minutes' ),
('event 6', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '2 hour' ),
('event 7', 'описание события','2020-04-01 09:00:00'::timestamp , '2020-04-01 10:30:00'::timestamp, 'User1', interval '1 day' )
;