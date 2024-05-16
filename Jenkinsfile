/* Requires the Docker Pipeline plugin */
pipeline {
    agent { docker { image 'golang:1.22.3-alpine3.19' } }
    stages {
        stage('build') {
            steps {
                sh 'echo "Hello World"'
                sh 'go build'
            }
        }
    }
}