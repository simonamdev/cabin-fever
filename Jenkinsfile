pipeline {
    agent none
    stages {
        stage('Build Frontend') {
            agent {
                docker { image 'node:14-alpine' }
            }
            steps {
                dir('cabinclient') {
                    sh 'yarn'
                    sh 'yarn run test'
                    sh 'yarn run build'
                    archiveArtifacts artifacts: 'dist/**/*.*'
                }
            }
        }
        stage('Build Backend') {
            agent {
                docker { image 'golang' }
            }
            steps {
                dir('cabinserver') {
                    // Add make to the docker container. TODO: Build this from a dockerifle?
                    sh 'apk update'
                    sh 'apk add make'
                    sh 'make test'
                    sh 'make build'
                    archiveArtifacts artifacts: 'cabinserver'
                }
            }
        }
    }
    
}
