pipeline { 
    agent { dockerfile true }
    stages {
        stage('Archive') {
            steps {
                archiveArtifacts artifacts: 'static/**/*.*'
                archiveArtifacts artifacts: 'cabinserver'
            }
        }
    }
    
}
