pipeline {
    options {
        buildDiscarder logRotator(artifactDaysToKeepStr: '5', artifactNumToKeepStr: '5', daysToKeepStr: '5', numToKeepStr: '5')
        durabilityHint 'PERFORMANCE_OPTIMIZED'
        timeout(5)
    }
    agent {
        kubernetes {
            label 'mypod'
            defaultContainer 'jnlp'
            yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    some-label: some-label-value
spec:
  containers:
  - name: go
    image: caladreas/go-build-agent
    command:
    - cat
    tty: true
"""
        }
    }
    stages {
        stage('Test versions') {
            steps {
                container('go') {
                    sh 'uname -a'
                    sh 'go version'
                }
            }
        }
        stage('Checkout') {
            steps {
                git 'https://github.com/joostvdg/go_demo_lib.git'
            }
        }
        stage('Build') {
            steps {
                container('go') {
                    sh 'go test --cover'
                }
            }
        }
    }
}