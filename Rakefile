
%w(
  jslint
  go
).each do |cmd|
  task develop: "which:#{cmd}"

  task "which:#{cmd}" do
    system("which #{cmd} > /dev/null") || fail("#{cmd} not found!")
  end
end


%w(
  github.com/antage/eventsource/http
  github.com/gorilla/mux
).each do |dep|
  dep_path = File.join(ENV['GOPATH'], dep)

  file dep_path do
    system("go get #{dep}") || fail
  end

  task 'go:get-dependencies' => dep_path
end


file TEDxMileHigh: 'go:get-dependencies' do
  system('go build') || fail
end


desc 'Start developing!'
task develop: :TEDxMileHigh do
  system('bundle && bundle exec guard') || fail
end
