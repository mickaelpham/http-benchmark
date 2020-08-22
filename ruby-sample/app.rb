# frozen_string_literal: true

require 'sinatra'

QUOTES = File.readlines('../quotes.txt')

get '/franklin-says' do
  [200, { 'Content-Type' => 'text/plain' }, QUOTES.sample]
end
