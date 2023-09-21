#! /usr/bin/env bash

kitex -module dousheng/publish -service publish ../dousheng_idl/publish/publish.thrift 
kitex -module dousheng/publish ../dousheng_idl/userinfo/userinfo.thrift 