pipeline {
    agent none
    stages {
        stage('Build Frontend') {
            agent {
                docker { image 'node' }
            }
            steps {
                dir('cabinclient') {
                    sh 'make install'
                    sh 'make test'
                    sh 'make build'
                    sh 'rm -rf static/'
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
                    archiveArtifacts artifacts: 'cabinserver'
                }
            }
        }
    }
    
}
