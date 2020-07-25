pipeline {
   agent any

   stages {
      stage('test') {
         steps {
            sh 'echo "test"'
            sh '''
                 make test
               '''
         }
      }
      stage('build') {
         steps {
            sh 'echo "build"'
            sh '''
                 make build
               '''
         }
      }
      stage('run') {
         steps {
            sh 'echo "run"'
            sh '''
                 make run
               '''
         }
      }
      stage('deploy') {
         steps {
            sh 'echo "deploy"'
         }
      }
   }
}
