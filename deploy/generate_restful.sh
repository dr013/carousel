#!/usr/bin/env bash
revel_mgo generate model Schema -fields=task:string,db_name:string,db_login:string,db_host:string,db_password:string,db_port:int,db_size:int,locked_date:datetime
revel_mgo generate controller Schema
