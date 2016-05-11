stage 'Commit'

node {
  def workspace = pwd()
  env.GOPATH = "${workspace}"
  echo "${env.BRANCH_NAME}"
  checkout scm
  sh 'go test -v github.com/superboum/atuin/...'
}
