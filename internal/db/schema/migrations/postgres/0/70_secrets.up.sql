begin;
/*
                                   ┌─────────────────┐
┌─────────────────┐                │ secrets_session │
│     secrets     │                ├─────────────────┤
├─────────────────┤                │public_id        │
│public_id        │┼─────────────○┼│scope_id         │
│name             │                │auth_token_id    │
│code             │                │user_id          │
│scope_id         │                │description      │
└─────────────────┘                └─────────────────┘



*/


-- secrets_scope_valid() is a before insert trigger function for secrets
create or replace function
  secrets_scope_valid()
  returns trigger
as $$
declare scope_type text;
begin
  -- Fetch the type of scope
  select isc.type from iam_scope isc where isc.public_id = new.scope_id into scope_type;
  if scope_type = 'project' then
    return new;
  end if;
  raise exception 'invalid secrets scope type % (must be project)', scope_type;
end;
$$ language plpgsql;



create table secrets (
  public_id wt_public_id primary key,
  name text not null,
  manager text not null,
  code text not null,
  scope_id wt_scope_id not null 
    references iam_scope(public_id) 
    on update cascade,
  create_time wt_timestamp,
  update_time wt_timestamp
);

create trigger 
  immutable_columns
before
update on secrets
  for each row execute procedure immutable_columns('public_id', 'scope_id', 'create_time');

create trigger 
  default_create_time_column
before
insert on secrets
  for each row execute procedure default_create_time();

create trigger 
  secrets_scope_valid
before insert on secrets
  for each row execute procedure secrets_scope_valid();


/** 
  Secrets Sessions
*/
create table secret_sessions (
    public_id wt_public_id primary key,
    description text,
    -- the secret
    secret_id text
      references secrets (public_id)
      on delete set null
      on update cascade, 
    -- the user of the session
    user_id text -- fk1
      -- not using the wt_user_id domain type because it is marked 'not null'
      references iam_user (public_id)
      on update cascade,
    -- the auth token of the user when this session was created
    -- auth_token_id wt_public_id -- fk6
    --   references auth_token (public_id)
    --   on delete set null
    --   on update cascade,
    -- the project which owns this session
    scope_id wt_scope_id -- fk7
      references iam_scope_project (scope_id)
      on update cascade,
    requested_time int not null,
    create_time wt_timestamp,
    updated_time wt_timestamp
);


create trigger 
  immutable_columns
before
update on secret_sessions
  for each row execute procedure immutable_columns('public_id', 'scope_id', 'create_time');

create trigger 
  default_create_time_column
before
insert on secret_sessions
  for each row execute procedure default_create_time();

create trigger 
  secrets_scope_valid
before insert on secret_sessions
  for each row execute procedure secrets_scope_valid();


/**
  Secrets Logs
*/

create table secret_session_logs (
    public_id wt_public_id primary key,
    secret_session_id text 
      references secret_sessions (public_id)
      on delete set null
      on update cascade, 
    output text,
    create_time wt_timestamp
);


create trigger 
  immutable_columns
before
update on secret_session_logs
  for each row execute procedure immutable_columns('public_id', 'create_time');

create trigger 
  default_create_time_column
before
insert on secret_session_logs
  for each row execute procedure default_create_time();


commit;