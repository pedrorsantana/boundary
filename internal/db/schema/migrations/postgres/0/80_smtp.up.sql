begin;
/*
                                   
┌─────────────────┐                ┌─────────────────┐
│     smtp_conf   │                │ iam_scope_org   │
├─────────────────┤                ├─────────────────┤
│public_id  (pk)  │┼──────────────┼│public_id  (pk)  │
│org_id     (fk)  │                │parent_id        │
│relay            │                │name             │
│port             │                └─────────────────┘
│sender           │                
│pass             │                
└─────────────────┘                
*/
/*
create table smtp_conf (
  public_id wt_public_id primary key,
  relay text not null,
  port int not null,
  sender text not null,
  pass text not null,
  org_id wt_scope_id not null 
    references iam_scope_org (public_id) 
    on delete cascade 
    on update cascade
);
*/
commit;