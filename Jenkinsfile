node {

   stage 'Checkout'
   env.GOPATH = "${env.PWD}/go"
   echo "${env.WORKSPACE}/go"
   sh '''
   env
   go get github.com/adamwalach/go-scroll-btn-demo
   '''
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
