node {
   env.WORKSPACE = pwd()
   env.GOPATH="${env.WORKSPACE}/go"
   env.GOBIN="${env.WORKSPACE}/go/bin"

   stage 'Check environment'
   echo """
         WORKSPACE: ${env.WORKSPACE}
         GOPATH: ${env.GOPATH}
         GOBIN: ${env.GOBIN}
   """

   stage 'Checkout'
   env.GOPATH = "${env.WORKSPACE}/go"
   echo "${env.PWD}"
   echo "Workspace: ${env.WORKSPACE}/go"
   sh '''
     export GOPATH="$WORKSPACE/go"
     echo $GOPATH
     mkdir -p $GOPATH
     go get -u "github.com/adamwalach/go-scroll-btn-demo"
   '''
   echo "${env.PWD}"
   //checkout scm

   stage 'Project build'
   sh '''
     export GOPATH="$WORKSPACE/go"
     cd $GOPATH/src/github.com/adamwalach/go-scroll-btn-demo
     /usr/bin/go version
     go build -o main *.go
   '''

   stage 'Docker build'
     sh '''#!/bin/bash
      echo $GOPATH
     export GOPATH="$WORKSPACE/go"
     docker build -f $GOPATH/src/github.com/adamwalach/go-scroll-btn-demo/Dockerfile \
                  -t awalach/go-scroll-btn-demo:$BRANCH_NAME ./
   '''

   stage 'Docker push'
     sh '''#!/bin/bash
     docker push awalach/go-scroll-btn-demo:$BRANCH_NAME ./
   '''
}
