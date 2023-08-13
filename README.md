<h1>Golang에서의 ElasticSearch 코드 작성</h1>

<h6>API를 연동해서, CRUD작업</h6>
- 메인 DB를 Mysql, Mongo가 아닌 ElasticSearch로 구성하여

<h6>해당 코드에서는 이론적인 내용은 다루지 않고, 코드로써는 어떻게 동작시키는지에 대해서 주롤 다룰 예정</h6>

- 부가적으로 메서드의 역할에 대해서는 다룰 수 있습니다.


<h3>Docker</h3>
```
docker run -d -p 9200:9200 -p 9300:9300 \
-e "discovery.type=single-node" \
-e "ELASTIC_USERNAME=<사용할 이름>" \
-e "ELASTIC_PASSWORD=<사용할 패스워드>" \
--name elasticsearch-docker \
docker.elastic.co/elasticsearch/elasticsearch:7.14.0
```