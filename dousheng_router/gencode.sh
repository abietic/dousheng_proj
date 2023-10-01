#!/bin/bash


hz new -I ../dousheng_idl/router/ -idl ../dousheng_idl/router/douyin/core.proto -module dousheng/router
hz update -I ../dousheng_idl/router/ -idl ../dousheng_idl/router/douyin/first.proto
hz update -I ../dousheng_idl/router/ -idl ../dousheng_idl/router/douyin/second.proto

