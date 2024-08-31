pipeline {
    agent any

    stages {
        stage('Pull Code') {
            steps {
                git branch: 'main', credentialsId: 'gitlab_auth', url: 'http://192.168.246.152:8081/devops/blog.git'
            }
        }
        stage('Build And Push Docker Image') {
            steps {
                sh 'docker build -t go-blog:1.1 .'
                withCredentials([usernamePassword(credentialsId: 'harbor_auth', passwordVariable: 'password', usernameVariable: 'username')]) {
                    sh 'docker login --username=$username --password=$password 192.168.246.152'
                    sh 'docker tag go-blog:1.1 192.168.246.152/blog/go-blog:1.1'
                    sh 'docker push 192.168.246.152/blog/go-blog:1.1'
                }
            }
        }
        stage('Deploy Remote Server') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'harbor_auth', passwordVariable: 'password', usernameVariable: 'username')]) {
                    sshPublisher(publishers: [sshPublisherDesc(configName: 'deploy_server', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: 'sh ./script/deploy.sh $password', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: 'script', remoteDirectorySDF: false, removePrefix: '', sourceFiles: 'deploy.sh')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
                }
            }
        }
    }
}
