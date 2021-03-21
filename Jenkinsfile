pipeline {
    agent none
    stages {
        stage('Build Frontend') {
            agent {
                docker { image 'node' }
            }
            steps {
                dir('cabinclient') {
                    sh 'yarn'
                    sh 'yarn run test'
                    sh 'yarn run build'
                    sh 'mv dist/ static/'
                    stash(name: 'frontend', includes: 'static/**/*')
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
                    sh 'make test'
                    unstash('frontend')
                    sh 'make build'
                }
            }
        }
    }
    
}
