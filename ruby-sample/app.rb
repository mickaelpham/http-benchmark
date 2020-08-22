# frozen_string_literal: true

require 'sinatra'

set :logging, false

QUOTES = File.readlines('../quotes.txt').freeze

get '/franklin-says' do
  [200, { 'Content-Type' => 'text/plain' }, QUOTES.sample]
end
