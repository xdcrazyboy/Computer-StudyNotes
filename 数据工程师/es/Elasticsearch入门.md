

## 如何在开发机上运行多个Elasticsearch实例
bin/elasticsearch -E node.name=node0 -E cluster.name=clusterDemo -E path.data=node0_data -d
bin/elasticsearch -E node.name=node1 -E cluster.name=clusterDemo -E path.data=node1_data -d
bin/elasticsearch -E node.name=node2 -E cluster.name=clusterDemo -E path.data=node2_data -d
bin/elasticsearch -E node.name=node3 -E cluster.name=clusterDemo -E path.data=node4_data -d