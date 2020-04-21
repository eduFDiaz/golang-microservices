pipeline {
   agent any

   stages {
      stage('git') {
         steps {
            echo 'Hello git'
            git test
         }
      }
      stage('test') {
         steps {
            echo 'Hello test'
            make test
         }
      }
      stage('build') {
         steps {
            echo 'Hello build'
            make build
         }
      }
      stage('run') {
         steps {
            echo 'Hello run'
            make run
         }
      }
      stage('deploy') {
         steps {
            echo 'Hello deploy'
         }
      }
   }
}
