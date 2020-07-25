pipeline {
   agent any

   stages {
      stage('test') {
         steps {
            sh -c '''echo 'Hello test'
                     make test'''
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
