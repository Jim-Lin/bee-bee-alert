# https://github.com/elastic/elasticsearch-docker
FROM docker.elastic.co/elasticsearch/elasticsearch-oss:6.0.0

# Add your elasticsearch plugins setup here
# Example: RUN elasticsearch-plugin install analysis-icu

COPY ./config/elasticsearch.yml /usr/share/elasticsearch/config/
COPY ./config/stopwords /usr/share/elasticsearch/config/stopwords
COPY ./plugins /usr/share/elasticsearch/plugins
