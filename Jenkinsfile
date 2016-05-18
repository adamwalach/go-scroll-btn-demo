node {
   env.WORKSPACE = pwd()
   env.GOPATH="${env.WORKSPACE}/go"
   env.GOBIN="${env.WORKSPACE}/go/bin"

   env.PROJECT_NAME="adamwalach/go-scroll-btn-demo"
   env.PROJECT_URL="github.com/${env.PROJECT_NAME}"
   env.PROJECT_PATH="${env.GOPATH}/src/${env.PROJECT_URL}"

   env.IMAGE_NAME="awalach/go-scroll-btn-demo"

   //for ansible (https://issues.jenkins-ci.org/browse/JENKINS-32911):
   env.PYTHONBUFFERED=0

   hipchatSend "Build ${env.BUILD_NUMBER} started. Project: '${env.PROJECT_NAME}', branch: '${env.BRANCH_NAME}'"
   stage 'Check environment'
     echo """
       WORKSPACE: ${env.WORKSPACE}
       GOPATH: ${env.GOPATH}
       GOBIN: ${env.GOBIN}

       PROJECT_NAME: ${env.PROJECT_NAME}
       PROJECT_URL: ${env.PROJECT_URL}
       PROJECT_PATH: ${env.PROJECT_PATH}
     """

   stage 'Cleanup'
     deleteDir()

   stage 'Checkout'
     sh '''
       mkdir -p "$PROJECT_PATH"
     '''
     dir ("${env.PROJECT_PATH}") {
       checkout scm
     }

   stage 'Tests'
     dir ("${env.PROJECT_PATH}") {
       sh '''
         gometalinter --vendor --fast --disable=gotype
       '''
     }

   stage 'Project build'
     dir ("${env.PROJECT_PATH}") {
       sh '''
         ./build.sh
       '''

       archive([includes: 'main'])
     }

   stage 'Docker build'
     dir ("${env.PROJECT_PATH}") {
       sh '''
         docker build -t $IMAGE_NAME:$BRANCH_NAME ./
       '''
     }

   stage 'Docker push'
     dir ("${env.PROJECT_PATH}") {
       sh '''#!/bin/bash
         time docker push $IMAGE_NAME:$BRANCH_NAME
       '''
   }

   stage 'Deploy'
     dir ("${env.PROJECT_PATH}") {
       ansiblePlaybook([playbook: 'playbook.yml'])
     }
     hipchatSend "Build ${env.BUILD_NUMBER} successfully finished. Project: '${env.PROJECT_NAME}', branch: '${env.BRANCH_NAME}'"
}
