pipeline{
    agent{
        label 'linux-ansible'
    }
    environment{
        DOCKERHUB_CREDENTIALS=credentials('karma792-dockerhub')
    }
    stages{
        stage('clone from git') {
            steps {
                git url:'https://github.com/ankan792/url-shortening-service-GO.git', branch: 'main'
            }
        }
        stage('Build docker image'){
            steps{
                sh '''
                    docker build -t karma792/url-shortener:$BUILD_NUMBER ./api/
                '''
            }
        }
        stage('Dockerhub login'){
            steps{
                sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
            }
        }
        stage('Push docker image'){
            steps{
                sh 'docker image push karma792/url-shortener:$BUILD_NUMBER'
            }
        }

        stage('Ansible'){
            steps{
                sh ''' 
                    ansible --version 
                    ansible-galaxy collection install community.docker
                '''
            }
        }
        stage('ansible-install docker on target'){
            steps{
                sh 'ansible-playbook ansible/playbooks/playbook-docker-install.yml'
            }
        }
        stage('ansible-run webapp'){
            steps{
                sh 'ansible-playbook ansible/playbooks/playbook-mywebapp-container.yml'
            }
        }

    }
}
