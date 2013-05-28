
{
  'jslint' => 'npm install jslint -g',
  'go'     => 'brew install go',
}.each do |cmd, installer|
  task develop: "which:#{cmd}"

  desc "locates #{cmd} on your PATH or complains loudly"
  task "which:#{cmd}" do
    system("which #{cmd} > /dev/null") || fail("#{cmd} not found!\nHave you installed it?\nTry: `#{installer}`")
  end
end


%w(
  github.com/antage/eventsource/http
  github.com/gorilla/mux
  github.com/laurent22/toml-go/toml
).each do |dep|
  dep_path = File.join(ENV['GOPATH'], dep)

  file dep_path do
    system("go get #{dep}") || fail
  end

  task 'go:dependencies' => dep_path
end


file TEDxMileHigh: 'go:dependencies' do
  system('go build') || fail
end


namespace :go do
	desc 'go get declared dependencies'
	task :dependencies
end


desc 'Start developing!'
task develop: :TEDxMileHigh do
  system('bundle && bundle exec guard') || fail
end

task default: :develop
