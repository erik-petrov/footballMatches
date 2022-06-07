create table users(
  id int primary key generated always as identity,
  name varchar(50) not null,
  email varchar(100) not null,
  password varchar(100) not null,
  created_at      timestamp       not null default now(),
  updated_at      timestamp       not null default now(),
  birthday date,
  gender text
);

create table teams(
  id int primary key generated always as identity,
  name char(50) not null,
  coach char(100)
);

create table matches(
  id int primary key generated always as identity,
  date date not null default now(),
  length time not null,
  team1 int,
  team2 int,
  score char(7),
  inProgress boolean,
  foreign key (team1)
    references teams(id),
  foreign key (team2)
    references teams(id)
);

insert into teams(name, coach) values('Liverpool', 'Jürgen Norbert Klopp');
insert into teams(name, coach) values('Manchester City', 'Josep Guardiola i Sala');
insert into teams(name, coach) values('Read Madrid', 'Carlo Ancelotti');

insert into matches(date, length, team1, team2, score, inProgress) values('2022-04-03', '01:05:39', 3, 2, '5:1', false);
insert into matches(date, length, team1, team2, score, inProgress) values(current_date, '01:10:02', 1, 3, '13:10', false);
insert into matches(date, length, team1, team2, score, inProgress) values(current_date, '01:11:12', 2, 1, '5:1', true);
insert into matches(date, length, team1, team2, score, inProgress) values('2022-07-04', null, 2, 3, null, false);