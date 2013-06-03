# vi: set ft=ruby :

interactor :off

group :services do

  guard :livereload do
    watch %r{^website/assets/css/.*\.css$}
    watch %r{^website/assets/ejs/.*\.ejs$}
    watch %r{^website/assets/img/.*\.(gif|jpg|png)$}
    watch %r{^website/assets/js/.*\.js$}
    watch %r{^tmp/website.pid$}
  end

  guard :process, name: 'compass' , command: 'compass watch --config website/assets/config.rb --quiet --boring'
  guard(:process, name: 'manager' , command: './TEDxMileHigh manager ') { watch(%r{^TEDxMileHigh$}); watch(%r{\.toml$}) }
  guard(:process, name: 'streamer', command: './TEDxMileHigh streamer') { watch(%r{^TEDxMileHigh$}); watch(%r{\.toml$}) }
  guard(:process, name: 'website' , command: './TEDxMileHigh website ') { watch(%r{^TEDxMileHigh$}); watch(%r{\.toml$}) }

end


group :build do

  guard :shell do
    # anytime a js file changes ...
    watch(%r{^website/assets/js/.*\.js$}) do |match|
      filename = match[0]
      system("jslint --color #{filename}")
    end
  end

  guard :shell do
    # anytime a go file changes ...
    watch(%r{.*\.go$}) do |m|
      next if m[0].start_with?('experiments')

      # announce we're going to build
      system('echo `date "+%H:%M:%S"` - building')

      # build our app; bail if it doesn't go well
      system('go build') || next

      # announce build passed and we're going to test
      system('echo `date "+%H:%M:%S"` - build OK - testing')

      # test our app; bail if it doesn't go well
      system('go test ./...') || next

      # announce test passed and we're going to vet
      system('echo `date "+%H:%M:%S"` - test OK - vetting')

      # vet our app; bail if it doesn't go well
      system('go vet ./...') || next

      # announce vet passed
      system('echo `date "+%H:%M:%S"` - vet OK')

    end
  end
end
