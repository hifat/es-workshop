# Elasticsearch Workshop

## Reference
- Mikelopster: [มาลองทำ Search ด้วย Elasticsearch กัน](https://www.youtube.com/watch?v=ECnEzu9E7Wg&t=2346s&ab_channel=mikelopster)
- CodeBangkok: [Elasticsearch: Get Started](https://www.youtube.com/watch?v=d2Ek-8HKdcI&list=PLyZTXfAT27iYpJgj--2fV7JDC7kYKYoUO&ab_channel=CodeBangkok)


## Issue & Solve

- ERROR: [1] bootstrap checks failed. You must address the points described in the following [1] lines before starting Elasticsearch. es01 | bootstrap check failure [1] of [1]: max virtual memory areas vm.max_map_count [65530] is too low ...

```docker
# Run in terminal
docker run -it --rm --privileged --pid=host alpine nsenter -t 1 -m -u -n -i sh -c "sysctl -w vm.max_map_count=262144"
```
