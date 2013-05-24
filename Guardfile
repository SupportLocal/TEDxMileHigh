# vi: set ft=ruby :

guard :process, name: 'compass', command: 'compass watch --config assets/config.rb --quiet --boring'

guard :process, name: 'TEDxMileHigh', command: './TEDxMileHigh' do
  watch %r{^TEDxMileHigh$}
end

guard :shell do

  # anytime a go file changes ...
  watch(%r{.*\.go$}) do

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

guard :livereload do
  watch %r{^assets/css}
  watch %r{^assets/img}
  watch %r{^assets/js}
  watch %r{^TEDxMileHigh\.pid$}
end
