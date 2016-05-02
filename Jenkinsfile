node {
   env.WORKSPACE = pwd()
   env.GOPATH="${env.WORKSPACE}/go"
   env.GOBIN="${env.WORKSPACE}/go/bin"

   env.PROJECT_NAME="adamwalach/go-scroll-btn-demo"
   env.PROJECT_URL="github.com/${env.PROJECT_NAME}"
   env.PROJECT_PATH="${env.GOPATH}/src/${env.PROJECT_URL}"

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

   stage 'Project build'
     dir ("${env.PROJECT_PATH}") {
       sh '''
         /usr/bin/go version
         go get ./
         go build -o main *.go
       '''
     }

   stage 'Docker build'
     dir ("${env.PROJECT_PATH}") {
       sh '''
         docker build -t awalach/go-scroll-btn-demo:$BRANCH_NAME ./
       '''
     }

   stage 'Docker push'
     sh '''
       docker push awalach/go-scroll-btn-demo:$BRANCH_NAME
     '''
}
