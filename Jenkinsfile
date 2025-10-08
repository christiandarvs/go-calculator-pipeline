pipeline {
    agent{docker {image 'golang:1.25.1-alpine3.22' }}

    stages{    
        stage('Checkout') {
            steps{
                echo "Checking out the source code"
                checkout "https://github.com/christiandarvs/pipeline-experiment.git"
            }
        }

        stage('Build') {
            steps{
                echo "Buidling the Go application..."
                sh '''
                go mod init go-calculator
                go mod tiny
                go build -o main main.go
                '''
            }
        }

        stage("Test") {
            steps{
                echo "Running tests..."
                sh "go test ./... -v"
            }
        }

        stage("Run") {
            steps{
                echo "Running the compiled binary..."
                sh "./main"
            }
        }
    }
}