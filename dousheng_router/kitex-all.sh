#! /usr/bin/env bash

kitex -module dousheng/router ../dousheng_idl/userinfo/userinfo.thrift 
kitex -module dousheng/router ../dousheng_idl/publish/publish.thrift 
kitex -module dousheng/router ../dousheng_idl/auth/auth.thrift 