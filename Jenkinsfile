pipeline {
    agent { docker { image 'golang:1.23-alpine' } }

    stages {
        stage('Checkout') {
            steps {
                echo "Checking out the source code"
                git branch: 'main', url: 'https://github.com/christiandarvs/pipeline-experiment.git'
            }
        }

        stage('Build') {
            steps {
                echo 'Building the Go application...'
                sh '''
                    go mod tidy
                    go build -o main main.go
                '''
            }
        }

        stage('Test') {
            steps {
                echo 'Running tests...'
                sh 'go test ./... -v'
            }
        }

        stage('Run') {
            steps {
                echo 'Running the compiled binary...'
                sh './main'
            }
        }
    }

    post {
        always {
            echo 'Pipeline completed.'
        }
        success {
            echo '✅ Build and tests passed successfully!'
        }
        failure {
            echo '❌ Build or tests failed.'
        }
    }
}
