pipeline {
   agent any
   triggers {
            pollSCM 'H/5 * * * *'
    }
    options {
        buildDiscarder logRotator(artifactDaysToKeepStr: '', artifactNumToKeepStr: '', daysToKeepStr: '', numToKeepStr: '5')
        timeout(time: 1, unit: 'DAYS')
    }

   stages {
         stage("Build") {
            steps {
                withCredentials([usernamePassword(credentialsId: 'aliyun', passwordVariable: 'password', usernameVariable: 'username')]) {
                        sh './ci upload ${username} ${password}'
                }

            }
        }


        stage('Deploy for SIT') {
            when {
                branch 'master'
            }
            steps {
                sh './ci deploy'
           }
        }

        stage('Deploy for UAT') {
            when {
                branch 'release'
            }
            steps {
                sh './ci deploy uat'
           }
        }
   }
  post {
    failure {
      emailext(
        subject: "FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
        body: """<p>FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]':</p>
                              <p>Check console output at "<a href="${env.BUILD_URL}">${env.JOB_NAME}          [${env.BUILD_NUMBER}]</a>"</p>""",
        to: "445999306@qq.com",
      )
    }
  }
}
