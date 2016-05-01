node {
   echo "${env.PWD}"
   stage 'Checkout'
   env.GOPATH = "${env.PWD}/go"
   echo "${env.PWD}"
   echo "${env.WORKSPACE}/go"
   sh '''
   echo "${env.PWD}"
   go get github.com/adamwalach/go-scroll-btn-demo
   '''
   echo "${env.PWD}"
   //checkout scm

   stage 'Project build'
   sh '''
   /usr/bin/go version
   go build -o main *.go
   '''

   stage 'Docker build'
   sh '''#!/bin/bash
   docker build -t awalach/go-scroll-btn-demo:$BRANCH_NAME ./
   '''

   stage 'Docker build'
   sh '''#!/bin/bash
   docker push awalach/go-scroll-btn-demo:$BRANCH_NAME ./
   '''
}
