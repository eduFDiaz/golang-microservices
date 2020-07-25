pipeline {
   agent any

   stages {
      stage('test') {
         steps {
            make test
         }
      }
      stage('build') {
         steps {
            make build
         }
      }
      stage('run') {
         steps {
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
