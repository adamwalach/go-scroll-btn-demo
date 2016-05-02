node {
   env.WORKSPACE = pwd()
   env.GOPATH="${env.WORKSPACE}/go"
   env.GOBIN="${env.WORKSPACE}/go/bin"

   env.PROJECT_NAME="adamwalach/go-scroll-btn-demo"
   env.PROJECT_URL="github.com/${env.PROJECT_NAME}"
   env.PROJECT_PATH="${env.GOPATH}/src/${env.PROJECT_URL}"

   env.IMAGE_NAME="awalach/go-scroll-btn-demo"

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
         GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=0 go build -o=main *.go
       '''
     }
   stage 'Tests'
     dir ("${env.PROJECT_PATH}") {
       sh '''
         gometalinter --fast
       '''
     }

   stage 'Docker build'
     dir ("${env.PROJECT_PATH}") {
       sh '''
         docker build -t $IMAGE_NAME:$BRANCH_NAME ./
       '''
     }

   stage 'Docker push'
     sh '''
       docker push $IMAGE_NAME:$BRANCH_NAME
     '''

   stage 'Deploy'
     echo "Deploy"
}
