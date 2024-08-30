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
    }
}
