
$gopath = ENV['GOPATH']

$dependencies = %w(
	code.google.com/p/go.net/websocket
	github.com/antage/eventsource/http
	github.com/gorilla/mux
)

$dependencies.each do |dep|
	dep_path = File.join($gopath, dep)

	file dep_path do
		system("go get #{dep}") || fail
	end

	task 'go:get-dependencies' => dep_path
end


namespace :go do
	desc 'go get known dependencies'
	task 'get-dependencies'
end

task default: 'go:get-dependencies'
