pipeline {
   agent any

   stages {
      stage('test') {
         steps {
            sh -c 'make test'
         }
      }
      stage('build') {
         steps {
            sh -c 'make build'
         }
      }
      stage('run') {
         steps {
            sh -c 'make run'
         }
      }
      stage('deploy') {
         steps {
            echo 'Hello deploy'
         }
      }
   }
}
