stage 'Commit'

node {
  def workspace = pwd()
  env.GOPATH = "${workspace}"

  checkout([$class: 'GitSCM', branches: [[name: "origin/${env.BRANCH_NAME}"]], doGenerateSubmoduleConfigurations: false, extensions: [[$class: 'RelativeTargetDirectory', relativeTargetDir: 'src/github.com/superboum/atuin']], submoduleCfg: [], userRemoteConfigs: [[url: 'https://github.com/superboum/atuin.git']]])

  sh 'go test -v github.com/superboum/atuin/...'
}
