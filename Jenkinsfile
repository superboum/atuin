stage 'Commit'

node {
  def workspace = pwd()
  env.GOPATH = "${workspace}"
  checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [[$class: 'CloneOption', noTags: false, reference: '', shallow: true], [$class: 'RelativeTargetDirectory', relativeTargetDir: 'github.com/superboum/atuin']], submoduleCfg: [], userRemoteConfigs: [[url: 'https://github.com/superboum/atuin']]])
  sh 'go test -v github.com/superboum/atuin/...'
}
