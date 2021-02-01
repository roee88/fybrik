---
hide:
  - toc        # Hide table of contents
---

# API Reference

Packages:

- motion.m4d.ibm.com/v1alpha1
    - [BatchTransfer](#batchtransfer)
    - [StreamTransfer](#streamtransfer)
- app.m4d.ibm.com/v1alpha1
    - [Blueprint](#blueprint)
    - [M4DApplication](#m4dapplication)
    - [M4DBucket](#m4dbucket)
    - [M4DModule](#m4dmodule)
    - [Plotter](#plotter)

## motion.m4d.ibm.com/v1alpha1

Resource Types:

- [BatchTransfer](#batchtransfer)

- [StreamTransfer](#streamtransfer)




### BatchTransfer
<sup><sup>[↩ Parent](#motion.m4d.ibm.com/v1alpha1 )</sup></sup>






BatchTransfer is the Schema for the batchtransfers API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>motion.m4d.ibm.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>BatchTransfer</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#batchtransferspec">spec</a></b></td>
        <td>object</td>
        <td>BatchTransferSpec defines the state of a BatchTransfer. The state includes source/destination specification, a schedule and the means by which data movement is to be conducted. The means is given as a kubernetes job description. In addition, the state also contains a sketch of a transformation instruction. In future releases, the transformation description should be specified in a separate CRD.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#batchtransferstatus">status</a></b></td>
        <td>object</td>
        <td>BatchTransferStatus defines the observed state of BatchTransfer This includes a reference to the job that implements the movement as well as the last schedule time. What is missing: Extended status information such as: - number of records moved - technical meta-data</td>
        <td>false</td>
      </tr></tbody>
</table>


#### BatchTransfer.spec
<sup><sup>[↩ Parent](#batchtransfer)</sup></sup>



BatchTransferSpec defines the state of a BatchTransfer. The state includes source/destination specification, a schedule and the means by which data movement is to be conducted. The means is given as a kubernetes job description. In addition, the state also contains a sketch of a transformation instruction. In future releases, the transformation description should be specified in a separate CRD.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>failedJobHistoryLimit</b></td>
        <td>integer</td>
        <td>Maximal number of failed Kubernetes job objects that should be kept. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b>flowType</b></td>
        <td>enum</td>
        <td>Data flow type that specifies if this is a stream or a batch workflow [Batch Stream]</td>
        <td>false</td>
      </tr><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>Image that should be used for the actual batch job. This is usually a datamover image. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>Image pull policy that should be used for the actual job. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b>maxFailedRetries</b></td>
        <td>integer</td>
        <td>Maximal number of failed retries until the batch job should stop trying. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b>noFinalizer</b></td>
        <td>boolean</td>
        <td>If this batch job instance should have a finalizer or not. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b>readDataType</b></td>
        <td>enum</td>
        <td>Data type of the data that is read from source (log data or change data) [LogData ChangeData]</td>
        <td>false</td>
      </tr><tr>
        <td><b>schedule</b></td>
        <td>string</td>
        <td>Cron schedule if this BatchTransfer job should run on a regular schedule. Values are specified like cron job schedules. A good translation to human language can be found here https://crontab.guru/</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretProviderRole</b></td>
        <td>string</td>
        <td>Secret provider role that should be used for the actual job. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretProviderURL</b></td>
        <td>string</td>
        <td>Secret provider url that should be used for the actual job. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#batchtransferspecspark">spark</a></b></td>
        <td>object</td>
        <td>Optional Spark configuration for tuning</td>
        <td>false</td>
      </tr><tr>
        <td><b>successfulJobHistoryLimit</b></td>
        <td>integer</td>
        <td>Maximal number of successful Kubernetes job objects that should be kept. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b>suspend</b></td>
        <td>boolean</td>
        <td>If this batch job instance is run on a schedule the regular schedule can be suspended with this property. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#batchtransferspectransformationindex">transformation</a></b></td>
        <td>[]object</td>
        <td>Transformations to be applied to the source data before writing to destination</td>
        <td>false</td>
      </tr><tr>
        <td><b>writeDataType</b></td>
        <td>enum</td>
        <td>Data type of how the data should be written to the target (log data or change data) [LogData ChangeData]</td>
        <td>false</td>
      </tr><tr>
        <td><b>writeOperation</b></td>
        <td>enum</td>
        <td>Write operation that should be performed when writing (overwrite,append,update) Caution: Some write operations are only available for batch and some only for stream. [Overwrite Append Update]</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#batchtransferspecdestination">destination</a></b></td>
        <td>object</td>
        <td>Destination data store for this batch job</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#batchtransferspecsource">source</a></b></td>
        <td>object</td>
        <td>Source data store for this batch job</td>
        <td>true</td>
      </tr></tbody>
</table>


#### BatchTransfer.spec.spark
<sup><sup>[↩ Parent](#batchtransferspec)</sup></sup>



Optional Spark configuration for tuning

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>appName</b></td>
        <td>string</td>
        <td>Name of the transaction. Mainly used for debugging and lineage tracking.</td>
        <td>false</td>
      </tr><tr>
        <td><b>driverCores</b></td>
        <td>integer</td>
        <td>Number of cores that the driver should use</td>
        <td>false</td>
      </tr><tr>
        <td><b>driverMemory</b></td>
        <td>integer</td>
        <td>Memory that the driver should have</td>
        <td>false</td>
      </tr><tr>
        <td><b>executorCores</b></td>
        <td>integer</td>
        <td>Number of cores that each executor should have</td>
        <td>false</td>
      </tr><tr>
        <td><b>executorMemory</b></td>
        <td>string</td>
        <td>Memory that each executor should have</td>
        <td>false</td>
      </tr><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>Image to be used for executors</td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>Image pull policy to be used for executor</td>
        <td>false</td>
      </tr><tr>
        <td><b>numExecutors</b></td>
        <td>integer</td>
        <td>Number of executors to be started</td>
        <td>false</td>
      </tr><tr>
        <td><b>options</b></td>
        <td>map[string]string</td>
        <td>Additional options for Spark configuration.</td>
        <td>false</td>
      </tr><tr>
        <td><b>shufflePartitions</b></td>
        <td>integer</td>
        <td>Number of shuffle partitions for Spark</td>
        <td>false</td>
      </tr></tbody>
</table>


#### BatchTransfer.spec.transformation[index]
<sup><sup>[↩ Parent](#batchtransferspec)</sup></sup>



to be refined...

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>action</b></td>
        <td>enum</td>
        <td>Transformation action that should be performed. [RemoveColumns EncryptColumns DigestColumns RedactColumns SampleRows FilterRows]</td>
        <td>false</td>
      </tr><tr>
        <td><b>columns</b></td>
        <td>[]string</td>
        <td>Columns that are involved in this action. This property is optional as for some actions no columns have to be specified. E.g. filter is a row based transformation.</td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>Name of the transaction. Mainly used for debugging and lineage tracking.</td>
        <td>false</td>
      </tr><tr>
        <td><b>options</b></td>
        <td>map[string]string</td>
        <td>Additional options for this transformation.</td>
        <td>false</td>
      </tr></tbody>
</table>


#### BatchTransfer.spec.destination
<sup><sup>[↩ Parent](#batchtransferspec)</sup></sup>



Destination data store for this batch job

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#batchtransferspecdestinationcloudant">cloudant</a></b></td>
        <td>object</td>
        <td>IBM Cloudant. Needs cloudant legacy credentials.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#batchtransferspecdestinationdatabase">database</a></b></td>
        <td>object</td>
        <td>Database data store. For the moment only Db2 is supported.</td>
        <td>false</td>
      </tr><tr>
        <td><b>description</b></td>
        <td>string</td>
        <td>Description of the transfer in human readable form that is displayed in the kubectl get If not provided this will be filled in depending on the datastore that is specified.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#batchtransferspecdestinationkafka">kafka</a></b></td>
        <td>object</td>
        <td>Kafka data store. The supposed format within the given Kafka topic is a Confluent compatible format stored as Avro. A schema registry needs to be specified as well.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#batchtransferspecdestinations3">s3</a></b></td>
        <td>object</td>
        <td>An object store data store that is compatible with S3. This can be a COS bucket.</td>
        <td>false</td>
      </tr></tbody>
</table>


#### BatchTransfer.spec.destination.cloudant
<sup><sup>[↩ Parent](#batchtransferspecdestination)</sup></sup>



IBM Cloudant. Needs cloudant legacy credentials.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>password</b></td>
        <td>string</td>
        <td>Cloudant password. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>username</b></td>
        <td>string</td>
        <td>Cloudant user. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the user name/password are stored. If not specified user and password have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>database</b></td>
        <td>string</td>
        <td>Database to be read from/written to</td>
        <td>true</td>
      </tr><tr>
        <td><b>host</b></td>
        <td>string</td>
        <td>Host of cloudant instance</td>
        <td>true</td>
      </tr></tbody>
</table>


#### BatchTransfer.spec.destination.database
<sup><sup>[↩ Parent](#batchtransferspecdestination)</sup></sup>



Database data store. For the moment only Db2 is supported.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>password</b></td>
        <td>string</td>
        <td>Database password. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>user</b></td>
        <td>string</td>
        <td>Database user. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the user name/password are stored. If not specified user and password have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>db2URL</b></td>
        <td>string</td>
        <td>URL to Db2 instance in JDBC format Supported SSL certificates are currently certificates signed with IBM Intermediate CA or cloud signed certificates.</td>
        <td>true</td>
      </tr><tr>
        <td><b>table</b></td>
        <td>string</td>
        <td>Table to be read</td>
        <td>true</td>
      </tr></tbody>
</table>


#### BatchTransfer.spec.destination.kafka
<sup><sup>[↩ Parent](#batchtransferspecdestination)</sup></sup>



Kafka data store. The supposed format within the given Kafka topic is a Confluent compatible format stored as Avro. A schema registry needs to be specified as well.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>createSnapshot</b></td>
        <td>boolean</td>
        <td>If a snapshot should be created of the topic. Records in Kafka are stored as key-value pairs. Updates/Deletes for the same key are appended to the Kafka topic and the last value for a given key is the valid key in a Snapshot. When this property is true only the last value will be written. If the property is false all values will be written out. As a CDC example: If the property is true a valid snapshot of the log stream will be created. If the property is false the CDC stream will be dumped as is like a change log.</td>
        <td>false</td>
      </tr><tr>
        <td><b>dataFormat</b></td>
        <td>string</td>
        <td>Data format of the objects in S3. e.g. parquet or csv. Please refer to struct for allowed values.</td>
        <td>false</td>
      </tr><tr>
        <td><b>keyDeserializer</b></td>
        <td>string</td>
        <td>Deserializer to be used for the keys of the topic</td>
        <td>false</td>
      </tr><tr>
        <td><b>password</b></td>
        <td>string</td>
        <td>Kafka user password Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>saslMechanism</b></td>
        <td>string</td>
        <td>SASL Mechanism to be used (e.g. PLAIN or SCRAM-SHA-512) Default SCRAM-SHA-512 will be assumed if not specified</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>securityProtocol</b></td>
        <td>string</td>
        <td>Kafka security protocol one of (PLAINTEXT, SASL_PLAINTEXT, SASL_SSL, SSL) Default SASL_SSL will be assumed if not specified</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststore</b></td>
        <td>string</td>
        <td>A truststore or certificate encoded as base64. The format can be JKS or PKCS12. A truststore can be specified like this or in a predefined Kubernetes secret</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststoreLocation</b></td>
        <td>string</td>
        <td>SSL truststore location.</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststorePassword</b></td>
        <td>string</td>
        <td>SSL truststore password.</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststoreSecret</b></td>
        <td>string</td>
        <td>Kubernetes secret that contains the SSL truststore. The format can be JKS or PKCS12. A truststore can be specified like this or as</td>
        <td>false</td>
      </tr><tr>
        <td><b>user</b></td>
        <td>string</td>
        <td>Kafka user name. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>valueDeserializer</b></td>
        <td>string</td>
        <td>Deserializer to be used for the values of the topic</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the user name/password are stored. If not specified user and password have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>kafkaBrokers</b></td>
        <td>string</td>
        <td>Kafka broker URLs as a comma separated list.</td>
        <td>true</td>
      </tr><tr>
        <td><b>kafkaTopic</b></td>
        <td>string</td>
        <td>Kafka topic</td>
        <td>true</td>
      </tr><tr>
        <td><b>schemaRegistryURL</b></td>
        <td>string</td>
        <td>URL to the schema registry. The registry has to be Confluent schema registry compatible.</td>
        <td>true</td>
      </tr></tbody>
</table>


#### BatchTransfer.spec.destination.s3
<sup><sup>[↩ Parent](#batchtransferspecdestination)</sup></sup>



An object store data store that is compatible with S3. This can be a COS bucket.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>accessKey</b></td>
        <td>string</td>
        <td>Access key of the HMAC credentials that can access the given bucket. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>dataFormat</b></td>
        <td>string</td>
        <td>Data format of the objects in S3. e.g. parquet or csv. Please refer to struct for allowed values.</td>
        <td>false</td>
      </tr><tr>
        <td><b>partitionBy</b></td>
        <td>[]string</td>
        <td>Partition by partition (for target data stores) Defines the columns to partition the output by for a target data store.</td>
        <td>false</td>
      </tr><tr>
        <td><b>region</b></td>
        <td>string</td>
        <td>Region of S3 service</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretKey</b></td>
        <td>string</td>
        <td>Secret key of the HMAC credentials that can access the given bucket. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the accessKey/secretKey are stored. If not specified accessKey and secretKey have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>bucket</b></td>
        <td>string</td>
        <td>Bucket of S3 service</td>
        <td>true</td>
      </tr><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td>Endpoint of S3 service</td>
        <td>true</td>
      </tr><tr>
        <td><b>objectKey</b></td>
        <td>string</td>
        <td>Object key of the object in S3. This is used as a prefix! Thus all objects that have the given objectKey as prefix will be used as input!</td>
        <td>true</td>
      </tr></tbody>
</table>


#### BatchTransfer.spec.source
<sup><sup>[↩ Parent](#batchtransferspec)</sup></sup>



Source data store for this batch job

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#batchtransferspecsourcecloudant">cloudant</a></b></td>
        <td>object</td>
        <td>IBM Cloudant. Needs cloudant legacy credentials.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#batchtransferspecsourcedatabase">database</a></b></td>
        <td>object</td>
        <td>Database data store. For the moment only Db2 is supported.</td>
        <td>false</td>
      </tr><tr>
        <td><b>description</b></td>
        <td>string</td>
        <td>Description of the transfer in human readable form that is displayed in the kubectl get If not provided this will be filled in depending on the datastore that is specified.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#batchtransferspecsourcekafka">kafka</a></b></td>
        <td>object</td>
        <td>Kafka data store. The supposed format within the given Kafka topic is a Confluent compatible format stored as Avro. A schema registry needs to be specified as well.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#batchtransferspecsources3">s3</a></b></td>
        <td>object</td>
        <td>An object store data store that is compatible with S3. This can be a COS bucket.</td>
        <td>false</td>
      </tr></tbody>
</table>


#### BatchTransfer.spec.source.cloudant
<sup><sup>[↩ Parent](#batchtransferspecsource)</sup></sup>



IBM Cloudant. Needs cloudant legacy credentials.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>password</b></td>
        <td>string</td>
        <td>Cloudant password. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>username</b></td>
        <td>string</td>
        <td>Cloudant user. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the user name/password are stored. If not specified user and password have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>database</b></td>
        <td>string</td>
        <td>Database to be read from/written to</td>
        <td>true</td>
      </tr><tr>
        <td><b>host</b></td>
        <td>string</td>
        <td>Host of cloudant instance</td>
        <td>true</td>
      </tr></tbody>
</table>


#### BatchTransfer.spec.source.database
<sup><sup>[↩ Parent](#batchtransferspecsource)</sup></sup>



Database data store. For the moment only Db2 is supported.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>password</b></td>
        <td>string</td>
        <td>Database password. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>user</b></td>
        <td>string</td>
        <td>Database user. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the user name/password are stored. If not specified user and password have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>db2URL</b></td>
        <td>string</td>
        <td>URL to Db2 instance in JDBC format Supported SSL certificates are currently certificates signed with IBM Intermediate CA or cloud signed certificates.</td>
        <td>true</td>
      </tr><tr>
        <td><b>table</b></td>
        <td>string</td>
        <td>Table to be read</td>
        <td>true</td>
      </tr></tbody>
</table>


#### BatchTransfer.spec.source.kafka
<sup><sup>[↩ Parent](#batchtransferspecsource)</sup></sup>



Kafka data store. The supposed format within the given Kafka topic is a Confluent compatible format stored as Avro. A schema registry needs to be specified as well.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>createSnapshot</b></td>
        <td>boolean</td>
        <td>If a snapshot should be created of the topic. Records in Kafka are stored as key-value pairs. Updates/Deletes for the same key are appended to the Kafka topic and the last value for a given key is the valid key in a Snapshot. When this property is true only the last value will be written. If the property is false all values will be written out. As a CDC example: If the property is true a valid snapshot of the log stream will be created. If the property is false the CDC stream will be dumped as is like a change log.</td>
        <td>false</td>
      </tr><tr>
        <td><b>dataFormat</b></td>
        <td>string</td>
        <td>Data format of the objects in S3. e.g. parquet or csv. Please refer to struct for allowed values.</td>
        <td>false</td>
      </tr><tr>
        <td><b>keyDeserializer</b></td>
        <td>string</td>
        <td>Deserializer to be used for the keys of the topic</td>
        <td>false</td>
      </tr><tr>
        <td><b>password</b></td>
        <td>string</td>
        <td>Kafka user password Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>saslMechanism</b></td>
        <td>string</td>
        <td>SASL Mechanism to be used (e.g. PLAIN or SCRAM-SHA-512) Default SCRAM-SHA-512 will be assumed if not specified</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>securityProtocol</b></td>
        <td>string</td>
        <td>Kafka security protocol one of (PLAINTEXT, SASL_PLAINTEXT, SASL_SSL, SSL) Default SASL_SSL will be assumed if not specified</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststore</b></td>
        <td>string</td>
        <td>A truststore or certificate encoded as base64. The format can be JKS or PKCS12. A truststore can be specified like this or in a predefined Kubernetes secret</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststoreLocation</b></td>
        <td>string</td>
        <td>SSL truststore location.</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststorePassword</b></td>
        <td>string</td>
        <td>SSL truststore password.</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststoreSecret</b></td>
        <td>string</td>
        <td>Kubernetes secret that contains the SSL truststore. The format can be JKS or PKCS12. A truststore can be specified like this or as</td>
        <td>false</td>
      </tr><tr>
        <td><b>user</b></td>
        <td>string</td>
        <td>Kafka user name. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>valueDeserializer</b></td>
        <td>string</td>
        <td>Deserializer to be used for the values of the topic</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the user name/password are stored. If not specified user and password have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>kafkaBrokers</b></td>
        <td>string</td>
        <td>Kafka broker URLs as a comma separated list.</td>
        <td>true</td>
      </tr><tr>
        <td><b>kafkaTopic</b></td>
        <td>string</td>
        <td>Kafka topic</td>
        <td>true</td>
      </tr><tr>
        <td><b>schemaRegistryURL</b></td>
        <td>string</td>
        <td>URL to the schema registry. The registry has to be Confluent schema registry compatible.</td>
        <td>true</td>
      </tr></tbody>
</table>


#### BatchTransfer.spec.source.s3
<sup><sup>[↩ Parent](#batchtransferspecsource)</sup></sup>



An object store data store that is compatible with S3. This can be a COS bucket.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>accessKey</b></td>
        <td>string</td>
        <td>Access key of the HMAC credentials that can access the given bucket. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>dataFormat</b></td>
        <td>string</td>
        <td>Data format of the objects in S3. e.g. parquet or csv. Please refer to struct for allowed values.</td>
        <td>false</td>
      </tr><tr>
        <td><b>partitionBy</b></td>
        <td>[]string</td>
        <td>Partition by partition (for target data stores) Defines the columns to partition the output by for a target data store.</td>
        <td>false</td>
      </tr><tr>
        <td><b>region</b></td>
        <td>string</td>
        <td>Region of S3 service</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretKey</b></td>
        <td>string</td>
        <td>Secret key of the HMAC credentials that can access the given bucket. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the accessKey/secretKey are stored. If not specified accessKey and secretKey have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>bucket</b></td>
        <td>string</td>
        <td>Bucket of S3 service</td>
        <td>true</td>
      </tr><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td>Endpoint of S3 service</td>
        <td>true</td>
      </tr><tr>
        <td><b>objectKey</b></td>
        <td>string</td>
        <td>Object key of the object in S3. This is used as a prefix! Thus all objects that have the given objectKey as prefix will be used as input!</td>
        <td>true</td>
      </tr></tbody>
</table>


#### BatchTransfer.status
<sup><sup>[↩ Parent](#batchtransfer)</sup></sup>



BatchTransferStatus defines the observed state of BatchTransfer This includes a reference to the job that implements the movement as well as the last schedule time. What is missing: Extended status information such as: - number of records moved - technical meta-data

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#batchtransferstatusactive">active</a></b></td>
        <td>object</td>
        <td>A pointer to the currently running job (or nil)</td>
        <td>false</td>
      </tr><tr>
        <td><b>error</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#batchtransferstatuslastcompleted">lastCompleted</a></b></td>
        <td>object</td>
        <td>ObjectReference contains enough information to let you inspect or modify the referred object. --- New uses of this type are discouraged because of difficulty describing its usage when embedded in APIs.  1. Ignored fields.  It includes many fields which are not generally honored.  For instance, ResourceVersion and FieldPath are both very rarely valid in actual usage.  2. Invalid usage help.  It is impossible to add specific help for individual usage.  In most embedded usages, there are particular     restrictions like, "must refer only to types A and B" or "UID not honored" or "name must be restricted".     Those cannot be well described when embedded.  3. Inconsistent validation.  Because the usages are different, the validation rules are different by usage, which makes it hard for users to predict what will happen.  4. The fields are both imprecise and overly precise.  Kind is not a precise mapping to a URL. This can produce ambiguity     during interpretation and require a REST mapping.  In most cases, the dependency is on the group,resource tuple     and the version of the actual struct is irrelevant.  5. We cannot easily change it.  Because this type is embedded in many locations, updates to this type     will affect numerous schemas.  Don't make new APIs embed an underspecified API type they do not control. Instead of using this type, create a locally provided and used type that is well-focused on your reference. For example, ServiceReferences for admission registration: https://github.com/kubernetes/api/blob/release-1.17/admissionregistration/v1/types.go#L533 .</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#batchtransferstatuslastfailed">lastFailed</a></b></td>
        <td>object</td>
        <td>ObjectReference contains enough information to let you inspect or modify the referred object. --- New uses of this type are discouraged because of difficulty describing its usage when embedded in APIs.  1. Ignored fields.  It includes many fields which are not generally honored.  For instance, ResourceVersion and FieldPath are both very rarely valid in actual usage.  2. Invalid usage help.  It is impossible to add specific help for individual usage.  In most embedded usages, there are particular     restrictions like, "must refer only to types A and B" or "UID not honored" or "name must be restricted".     Those cannot be well described when embedded.  3. Inconsistent validation.  Because the usages are different, the validation rules are different by usage, which makes it hard for users to predict what will happen.  4. The fields are both imprecise and overly precise.  Kind is not a precise mapping to a URL. This can produce ambiguity     during interpretation and require a REST mapping.  In most cases, the dependency is on the group,resource tuple     and the version of the actual struct is irrelevant.  5. We cannot easily change it.  Because this type is embedded in many locations, updates to this type     will affect numerous schemas.  Don't make new APIs embed an underspecified API type they do not control. Instead of using this type, create a locally provided and used type that is well-focused on your reference. For example, ServiceReferences for admission registration: https://github.com/kubernetes/api/blob/release-1.17/admissionregistration/v1/types.go#L533 .</td>
        <td>false</td>
      </tr><tr>
        <td><b>lastRecordTime</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>lastScheduleTime</b></td>
        <td>string</td>
        <td>Information when was the last time the job was successfully scheduled.</td>
        <td>false</td>
      </tr><tr>
        <td><b>lastSuccessTime</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>numRecords</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>status</b></td>
        <td>enum</td>
        <td> [STARTING RUNNING SUCCEEDED FAILED]</td>
        <td>false</td>
      </tr></tbody>
</table>


#### BatchTransfer.status.active
<sup><sup>[↩ Parent](#batchtransferstatus)</sup></sup>



A pointer to the currently running job (or nil)

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>apiVersion</b></td>
        <td>string</td>
        <td>API version of the referent.</td>
        <td>false</td>
      </tr><tr>
        <td><b>fieldPath</b></td>
        <td>string</td>
        <td>If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.</td>
        <td>false</td>
      </tr><tr>
        <td><b>kind</b></td>
        <td>string</td>
        <td>Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names</td>
        <td>false</td>
      </tr><tr>
        <td><b>namespace</b></td>
        <td>string</td>
        <td>Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/</td>
        <td>false</td>
      </tr><tr>
        <td><b>resourceVersion</b></td>
        <td>string</td>
        <td>Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency</td>
        <td>false</td>
      </tr><tr>
        <td><b>uid</b></td>
        <td>string</td>
        <td>UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids</td>
        <td>false</td>
      </tr></tbody>
</table>


#### BatchTransfer.status.lastCompleted
<sup><sup>[↩ Parent](#batchtransferstatus)</sup></sup>



ObjectReference contains enough information to let you inspect or modify the referred object. --- New uses of this type are discouraged because of difficulty describing its usage when embedded in APIs.  1. Ignored fields.  It includes many fields which are not generally honored.  For instance, ResourceVersion and FieldPath are both very rarely valid in actual usage.  2. Invalid usage help.  It is impossible to add specific help for individual usage.  In most embedded usages, there are particular     restrictions like, "must refer only to types A and B" or "UID not honored" or "name must be restricted".     Those cannot be well described when embedded.  3. Inconsistent validation.  Because the usages are different, the validation rules are different by usage, which makes it hard for users to predict what will happen.  4. The fields are both imprecise and overly precise.  Kind is not a precise mapping to a URL. This can produce ambiguity     during interpretation and require a REST mapping.  In most cases, the dependency is on the group,resource tuple     and the version of the actual struct is irrelevant.  5. We cannot easily change it.  Because this type is embedded in many locations, updates to this type     will affect numerous schemas.  Don't make new APIs embed an underspecified API type they do not control. Instead of using this type, create a locally provided and used type that is well-focused on your reference. For example, ServiceReferences for admission registration: https://github.com/kubernetes/api/blob/release-1.17/admissionregistration/v1/types.go#L533 .

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>apiVersion</b></td>
        <td>string</td>
        <td>API version of the referent.</td>
        <td>false</td>
      </tr><tr>
        <td><b>fieldPath</b></td>
        <td>string</td>
        <td>If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.</td>
        <td>false</td>
      </tr><tr>
        <td><b>kind</b></td>
        <td>string</td>
        <td>Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names</td>
        <td>false</td>
      </tr><tr>
        <td><b>namespace</b></td>
        <td>string</td>
        <td>Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/</td>
        <td>false</td>
      </tr><tr>
        <td><b>resourceVersion</b></td>
        <td>string</td>
        <td>Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency</td>
        <td>false</td>
      </tr><tr>
        <td><b>uid</b></td>
        <td>string</td>
        <td>UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids</td>
        <td>false</td>
      </tr></tbody>
</table>


#### BatchTransfer.status.lastFailed
<sup><sup>[↩ Parent](#batchtransferstatus)</sup></sup>



ObjectReference contains enough information to let you inspect or modify the referred object. --- New uses of this type are discouraged because of difficulty describing its usage when embedded in APIs.  1. Ignored fields.  It includes many fields which are not generally honored.  For instance, ResourceVersion and FieldPath are both very rarely valid in actual usage.  2. Invalid usage help.  It is impossible to add specific help for individual usage.  In most embedded usages, there are particular     restrictions like, "must refer only to types A and B" or "UID not honored" or "name must be restricted".     Those cannot be well described when embedded.  3. Inconsistent validation.  Because the usages are different, the validation rules are different by usage, which makes it hard for users to predict what will happen.  4. The fields are both imprecise and overly precise.  Kind is not a precise mapping to a URL. This can produce ambiguity     during interpretation and require a REST mapping.  In most cases, the dependency is on the group,resource tuple     and the version of the actual struct is irrelevant.  5. We cannot easily change it.  Because this type is embedded in many locations, updates to this type     will affect numerous schemas.  Don't make new APIs embed an underspecified API type they do not control. Instead of using this type, create a locally provided and used type that is well-focused on your reference. For example, ServiceReferences for admission registration: https://github.com/kubernetes/api/blob/release-1.17/admissionregistration/v1/types.go#L533 .

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>apiVersion</b></td>
        <td>string</td>
        <td>API version of the referent.</td>
        <td>false</td>
      </tr><tr>
        <td><b>fieldPath</b></td>
        <td>string</td>
        <td>If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.</td>
        <td>false</td>
      </tr><tr>
        <td><b>kind</b></td>
        <td>string</td>
        <td>Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names</td>
        <td>false</td>
      </tr><tr>
        <td><b>namespace</b></td>
        <td>string</td>
        <td>Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/</td>
        <td>false</td>
      </tr><tr>
        <td><b>resourceVersion</b></td>
        <td>string</td>
        <td>Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency</td>
        <td>false</td>
      </tr><tr>
        <td><b>uid</b></td>
        <td>string</td>
        <td>UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids</td>
        <td>false</td>
      </tr></tbody>
</table>

### StreamTransfer
<sup><sup>[↩ Parent](#motion.m4d.ibm.com/v1alpha1 )</sup></sup>






StreamTransfer is the Schema for the streamtransfers API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>motion.m4d.ibm.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>StreamTransfer</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#streamtransferspec">spec</a></b></td>
        <td>object</td>
        <td>StreamTransferSpec defines the desired state of StreamTransfer</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#streamtransferstatus">status</a></b></td>
        <td>object</td>
        <td>StreamTransferStatus defines the observed state of StreamTransfer</td>
        <td>false</td>
      </tr></tbody>
</table>


#### StreamTransfer.spec
<sup><sup>[↩ Parent](#streamtransfer)</sup></sup>



StreamTransferSpec defines the desired state of StreamTransfer

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>flowType</b></td>
        <td>enum</td>
        <td>Data flow type that specifies if this is a stream or a batch workflow [Batch Stream]</td>
        <td>false</td>
      </tr><tr>
        <td><b>image</b></td>
        <td>string</td>
        <td>Image that should be used for the actual batch job. This is usually a datamover image. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b>imagePullPolicy</b></td>
        <td>string</td>
        <td>Image pull policy that should be used for the actual job. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b>noFinalizer</b></td>
        <td>boolean</td>
        <td>If this batch job instance should have a finalizer or not. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b>readDataType</b></td>
        <td>enum</td>
        <td>Data type of the data that is read from source (log data or change data) [LogData ChangeData]</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretProviderRole</b></td>
        <td>string</td>
        <td>Secret provider role that should be used for the actual job. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretProviderURL</b></td>
        <td>string</td>
        <td>Secret provider url that should be used for the actual job. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b>suspend</b></td>
        <td>boolean</td>
        <td>If this batch job instance is run on a schedule the regular schedule can be suspended with this property. This property will be defaulted by the webhook if not set.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#streamtransferspectransformationindex">transformation</a></b></td>
        <td>[]object</td>
        <td>Transformations to be applied to the source data before writing to destination</td>
        <td>false</td>
      </tr><tr>
        <td><b>triggerInterval</b></td>
        <td>string</td>
        <td>Interval in which the Micro batches of this stream should be triggered The default is '5 seconds'.</td>
        <td>false</td>
      </tr><tr>
        <td><b>writeDataType</b></td>
        <td>enum</td>
        <td>Data type of how the data should be written to the target (log data or change data) [LogData ChangeData]</td>
        <td>false</td>
      </tr><tr>
        <td><b>writeOperation</b></td>
        <td>enum</td>
        <td>Write operation that should be performed when writing (overwrite,append,update) Caution: Some write operations are only available for batch and some only for stream. [Overwrite Append Update]</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#streamtransferspecdestination">destination</a></b></td>
        <td>object</td>
        <td>Destination data store for this batch job</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#streamtransferspecsource">source</a></b></td>
        <td>object</td>
        <td>Source data store for this batch job</td>
        <td>true</td>
      </tr></tbody>
</table>


#### StreamTransfer.spec.transformation[index]
<sup><sup>[↩ Parent](#streamtransferspec)</sup></sup>



to be refined...

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>action</b></td>
        <td>enum</td>
        <td>Transformation action that should be performed. [RemoveColumns EncryptColumns DigestColumns RedactColumns SampleRows FilterRows]</td>
        <td>false</td>
      </tr><tr>
        <td><b>columns</b></td>
        <td>[]string</td>
        <td>Columns that are involved in this action. This property is optional as for some actions no columns have to be specified. E.g. filter is a row based transformation.</td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>Name of the transaction. Mainly used for debugging and lineage tracking.</td>
        <td>false</td>
      </tr><tr>
        <td><b>options</b></td>
        <td>map[string]string</td>
        <td>Additional options for this transformation.</td>
        <td>false</td>
      </tr></tbody>
</table>


#### StreamTransfer.spec.destination
<sup><sup>[↩ Parent](#streamtransferspec)</sup></sup>



Destination data store for this batch job

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#streamtransferspecdestinationcloudant">cloudant</a></b></td>
        <td>object</td>
        <td>IBM Cloudant. Needs cloudant legacy credentials.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#streamtransferspecdestinationdatabase">database</a></b></td>
        <td>object</td>
        <td>Database data store. For the moment only Db2 is supported.</td>
        <td>false</td>
      </tr><tr>
        <td><b>description</b></td>
        <td>string</td>
        <td>Description of the transfer in human readable form that is displayed in the kubectl get If not provided this will be filled in depending on the datastore that is specified.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#streamtransferspecdestinationkafka">kafka</a></b></td>
        <td>object</td>
        <td>Kafka data store. The supposed format within the given Kafka topic is a Confluent compatible format stored as Avro. A schema registry needs to be specified as well.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#streamtransferspecdestinations3">s3</a></b></td>
        <td>object</td>
        <td>An object store data store that is compatible with S3. This can be a COS bucket.</td>
        <td>false</td>
      </tr></tbody>
</table>


#### StreamTransfer.spec.destination.cloudant
<sup><sup>[↩ Parent](#streamtransferspecdestination)</sup></sup>



IBM Cloudant. Needs cloudant legacy credentials.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>password</b></td>
        <td>string</td>
        <td>Cloudant password. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>username</b></td>
        <td>string</td>
        <td>Cloudant user. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the user name/password are stored. If not specified user and password have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>database</b></td>
        <td>string</td>
        <td>Database to be read from/written to</td>
        <td>true</td>
      </tr><tr>
        <td><b>host</b></td>
        <td>string</td>
        <td>Host of cloudant instance</td>
        <td>true</td>
      </tr></tbody>
</table>


#### StreamTransfer.spec.destination.database
<sup><sup>[↩ Parent](#streamtransferspecdestination)</sup></sup>



Database data store. For the moment only Db2 is supported.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>password</b></td>
        <td>string</td>
        <td>Database password. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>user</b></td>
        <td>string</td>
        <td>Database user. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the user name/password are stored. If not specified user and password have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>db2URL</b></td>
        <td>string</td>
        <td>URL to Db2 instance in JDBC format Supported SSL certificates are currently certificates signed with IBM Intermediate CA or cloud signed certificates.</td>
        <td>true</td>
      </tr><tr>
        <td><b>table</b></td>
        <td>string</td>
        <td>Table to be read</td>
        <td>true</td>
      </tr></tbody>
</table>


#### StreamTransfer.spec.destination.kafka
<sup><sup>[↩ Parent](#streamtransferspecdestination)</sup></sup>



Kafka data store. The supposed format within the given Kafka topic is a Confluent compatible format stored as Avro. A schema registry needs to be specified as well.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>createSnapshot</b></td>
        <td>boolean</td>
        <td>If a snapshot should be created of the topic. Records in Kafka are stored as key-value pairs. Updates/Deletes for the same key are appended to the Kafka topic and the last value for a given key is the valid key in a Snapshot. When this property is true only the last value will be written. If the property is false all values will be written out. As a CDC example: If the property is true a valid snapshot of the log stream will be created. If the property is false the CDC stream will be dumped as is like a change log.</td>
        <td>false</td>
      </tr><tr>
        <td><b>dataFormat</b></td>
        <td>string</td>
        <td>Data format of the objects in S3. e.g. parquet or csv. Please refer to struct for allowed values.</td>
        <td>false</td>
      </tr><tr>
        <td><b>keyDeserializer</b></td>
        <td>string</td>
        <td>Deserializer to be used for the keys of the topic</td>
        <td>false</td>
      </tr><tr>
        <td><b>password</b></td>
        <td>string</td>
        <td>Kafka user password Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>saslMechanism</b></td>
        <td>string</td>
        <td>SASL Mechanism to be used (e.g. PLAIN or SCRAM-SHA-512) Default SCRAM-SHA-512 will be assumed if not specified</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>securityProtocol</b></td>
        <td>string</td>
        <td>Kafka security protocol one of (PLAINTEXT, SASL_PLAINTEXT, SASL_SSL, SSL) Default SASL_SSL will be assumed if not specified</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststore</b></td>
        <td>string</td>
        <td>A truststore or certificate encoded as base64. The format can be JKS or PKCS12. A truststore can be specified like this or in a predefined Kubernetes secret</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststoreLocation</b></td>
        <td>string</td>
        <td>SSL truststore location.</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststorePassword</b></td>
        <td>string</td>
        <td>SSL truststore password.</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststoreSecret</b></td>
        <td>string</td>
        <td>Kubernetes secret that contains the SSL truststore. The format can be JKS or PKCS12. A truststore can be specified like this or as</td>
        <td>false</td>
      </tr><tr>
        <td><b>user</b></td>
        <td>string</td>
        <td>Kafka user name. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>valueDeserializer</b></td>
        <td>string</td>
        <td>Deserializer to be used for the values of the topic</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the user name/password are stored. If not specified user and password have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>kafkaBrokers</b></td>
        <td>string</td>
        <td>Kafka broker URLs as a comma separated list.</td>
        <td>true</td>
      </tr><tr>
        <td><b>kafkaTopic</b></td>
        <td>string</td>
        <td>Kafka topic</td>
        <td>true</td>
      </tr><tr>
        <td><b>schemaRegistryURL</b></td>
        <td>string</td>
        <td>URL to the schema registry. The registry has to be Confluent schema registry compatible.</td>
        <td>true</td>
      </tr></tbody>
</table>


#### StreamTransfer.spec.destination.s3
<sup><sup>[↩ Parent](#streamtransferspecdestination)</sup></sup>



An object store data store that is compatible with S3. This can be a COS bucket.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>accessKey</b></td>
        <td>string</td>
        <td>Access key of the HMAC credentials that can access the given bucket. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>dataFormat</b></td>
        <td>string</td>
        <td>Data format of the objects in S3. e.g. parquet or csv. Please refer to struct for allowed values.</td>
        <td>false</td>
      </tr><tr>
        <td><b>partitionBy</b></td>
        <td>[]string</td>
        <td>Partition by partition (for target data stores) Defines the columns to partition the output by for a target data store.</td>
        <td>false</td>
      </tr><tr>
        <td><b>region</b></td>
        <td>string</td>
        <td>Region of S3 service</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretKey</b></td>
        <td>string</td>
        <td>Secret key of the HMAC credentials that can access the given bucket. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the accessKey/secretKey are stored. If not specified accessKey and secretKey have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>bucket</b></td>
        <td>string</td>
        <td>Bucket of S3 service</td>
        <td>true</td>
      </tr><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td>Endpoint of S3 service</td>
        <td>true</td>
      </tr><tr>
        <td><b>objectKey</b></td>
        <td>string</td>
        <td>Object key of the object in S3. This is used as a prefix! Thus all objects that have the given objectKey as prefix will be used as input!</td>
        <td>true</td>
      </tr></tbody>
</table>


#### StreamTransfer.spec.source
<sup><sup>[↩ Parent](#streamtransferspec)</sup></sup>



Source data store for this batch job

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#streamtransferspecsourcecloudant">cloudant</a></b></td>
        <td>object</td>
        <td>IBM Cloudant. Needs cloudant legacy credentials.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#streamtransferspecsourcedatabase">database</a></b></td>
        <td>object</td>
        <td>Database data store. For the moment only Db2 is supported.</td>
        <td>false</td>
      </tr><tr>
        <td><b>description</b></td>
        <td>string</td>
        <td>Description of the transfer in human readable form that is displayed in the kubectl get If not provided this will be filled in depending on the datastore that is specified.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#streamtransferspecsourcekafka">kafka</a></b></td>
        <td>object</td>
        <td>Kafka data store. The supposed format within the given Kafka topic is a Confluent compatible format stored as Avro. A schema registry needs to be specified as well.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#streamtransferspecsources3">s3</a></b></td>
        <td>object</td>
        <td>An object store data store that is compatible with S3. This can be a COS bucket.</td>
        <td>false</td>
      </tr></tbody>
</table>


#### StreamTransfer.spec.source.cloudant
<sup><sup>[↩ Parent](#streamtransferspecsource)</sup></sup>



IBM Cloudant. Needs cloudant legacy credentials.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>password</b></td>
        <td>string</td>
        <td>Cloudant password. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>username</b></td>
        <td>string</td>
        <td>Cloudant user. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the user name/password are stored. If not specified user and password have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>database</b></td>
        <td>string</td>
        <td>Database to be read from/written to</td>
        <td>true</td>
      </tr><tr>
        <td><b>host</b></td>
        <td>string</td>
        <td>Host of cloudant instance</td>
        <td>true</td>
      </tr></tbody>
</table>


#### StreamTransfer.spec.source.database
<sup><sup>[↩ Parent](#streamtransferspecsource)</sup></sup>



Database data store. For the moment only Db2 is supported.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>password</b></td>
        <td>string</td>
        <td>Database password. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>user</b></td>
        <td>string</td>
        <td>Database user. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the user name/password are stored. If not specified user and password have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>db2URL</b></td>
        <td>string</td>
        <td>URL to Db2 instance in JDBC format Supported SSL certificates are currently certificates signed with IBM Intermediate CA or cloud signed certificates.</td>
        <td>true</td>
      </tr><tr>
        <td><b>table</b></td>
        <td>string</td>
        <td>Table to be read</td>
        <td>true</td>
      </tr></tbody>
</table>


#### StreamTransfer.spec.source.kafka
<sup><sup>[↩ Parent](#streamtransferspecsource)</sup></sup>



Kafka data store. The supposed format within the given Kafka topic is a Confluent compatible format stored as Avro. A schema registry needs to be specified as well.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>createSnapshot</b></td>
        <td>boolean</td>
        <td>If a snapshot should be created of the topic. Records in Kafka are stored as key-value pairs. Updates/Deletes for the same key are appended to the Kafka topic and the last value for a given key is the valid key in a Snapshot. When this property is true only the last value will be written. If the property is false all values will be written out. As a CDC example: If the property is true a valid snapshot of the log stream will be created. If the property is false the CDC stream will be dumped as is like a change log.</td>
        <td>false</td>
      </tr><tr>
        <td><b>dataFormat</b></td>
        <td>string</td>
        <td>Data format of the objects in S3. e.g. parquet or csv. Please refer to struct for allowed values.</td>
        <td>false</td>
      </tr><tr>
        <td><b>keyDeserializer</b></td>
        <td>string</td>
        <td>Deserializer to be used for the keys of the topic</td>
        <td>false</td>
      </tr><tr>
        <td><b>password</b></td>
        <td>string</td>
        <td>Kafka user password Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>saslMechanism</b></td>
        <td>string</td>
        <td>SASL Mechanism to be used (e.g. PLAIN or SCRAM-SHA-512) Default SCRAM-SHA-512 will be assumed if not specified</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>securityProtocol</b></td>
        <td>string</td>
        <td>Kafka security protocol one of (PLAINTEXT, SASL_PLAINTEXT, SASL_SSL, SSL) Default SASL_SSL will be assumed if not specified</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststore</b></td>
        <td>string</td>
        <td>A truststore or certificate encoded as base64. The format can be JKS or PKCS12. A truststore can be specified like this or in a predefined Kubernetes secret</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststoreLocation</b></td>
        <td>string</td>
        <td>SSL truststore location.</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststorePassword</b></td>
        <td>string</td>
        <td>SSL truststore password.</td>
        <td>false</td>
      </tr><tr>
        <td><b>sslTruststoreSecret</b></td>
        <td>string</td>
        <td>Kubernetes secret that contains the SSL truststore. The format can be JKS or PKCS12. A truststore can be specified like this or as</td>
        <td>false</td>
      </tr><tr>
        <td><b>user</b></td>
        <td>string</td>
        <td>Kafka user name. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>valueDeserializer</b></td>
        <td>string</td>
        <td>Deserializer to be used for the values of the topic</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the user name/password are stored. If not specified user and password have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>kafkaBrokers</b></td>
        <td>string</td>
        <td>Kafka broker URLs as a comma separated list.</td>
        <td>true</td>
      </tr><tr>
        <td><b>kafkaTopic</b></td>
        <td>string</td>
        <td>Kafka topic</td>
        <td>true</td>
      </tr><tr>
        <td><b>schemaRegistryURL</b></td>
        <td>string</td>
        <td>URL to the schema registry. The registry has to be Confluent schema registry compatible.</td>
        <td>true</td>
      </tr></tbody>
</table>


#### StreamTransfer.spec.source.s3
<sup><sup>[↩ Parent](#streamtransferspecsource)</sup></sup>



An object store data store that is compatible with S3. This can be a COS bucket.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>accessKey</b></td>
        <td>string</td>
        <td>Access key of the HMAC credentials that can access the given bucket. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>dataFormat</b></td>
        <td>string</td>
        <td>Data format of the objects in S3. e.g. parquet or csv. Please refer to struct for allowed values.</td>
        <td>false</td>
      </tr><tr>
        <td><b>partitionBy</b></td>
        <td>[]string</td>
        <td>Partition by partition (for target data stores) Defines the columns to partition the output by for a target data store.</td>
        <td>false</td>
      </tr><tr>
        <td><b>region</b></td>
        <td>string</td>
        <td>Region of S3 service</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretImport</b></td>
        <td>string</td>
        <td>Define a secret import definition.</td>
        <td>false</td>
      </tr><tr>
        <td><b>secretKey</b></td>
        <td>string</td>
        <td>Secret key of the HMAC credentials that can access the given bucket. Can be retrieved from vault if specified in vaultPath parameter and is thus optional.</td>
        <td>false</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Vault path where the accessKey/secretKey are stored. If not specified accessKey and secretKey have to be specified!</td>
        <td>false</td>
      </tr><tr>
        <td><b>bucket</b></td>
        <td>string</td>
        <td>Bucket of S3 service</td>
        <td>true</td>
      </tr><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td>Endpoint of S3 service</td>
        <td>true</td>
      </tr><tr>
        <td><b>objectKey</b></td>
        <td>string</td>
        <td>Object key of the object in S3. This is used as a prefix! Thus all objects that have the given objectKey as prefix will be used as input!</td>
        <td>true</td>
      </tr></tbody>
</table>


#### StreamTransfer.status
<sup><sup>[↩ Parent](#streamtransfer)</sup></sup>



StreamTransferStatus defines the observed state of StreamTransfer

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#streamtransferstatusactive">active</a></b></td>
        <td>object</td>
        <td>A pointer to the currently running job (or nil)</td>
        <td>false</td>
      </tr><tr>
        <td><b>error</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>status</b></td>
        <td>enum</td>
        <td> [STARTING RUNNING STOPPED FAILING]</td>
        <td>false</td>
      </tr></tbody>
</table>


#### StreamTransfer.status.active
<sup><sup>[↩ Parent](#streamtransferstatus)</sup></sup>



A pointer to the currently running job (or nil)

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>apiVersion</b></td>
        <td>string</td>
        <td>API version of the referent.</td>
        <td>false</td>
      </tr><tr>
        <td><b>fieldPath</b></td>
        <td>string</td>
        <td>If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: "spec.containers{name}" (where "name" refers to the name of the container that triggered the event) or if no container name is specified "spec.containers[2]" (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.</td>
        <td>false</td>
      </tr><tr>
        <td><b>kind</b></td>
        <td>string</td>
        <td>Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names</td>
        <td>false</td>
      </tr><tr>
        <td><b>namespace</b></td>
        <td>string</td>
        <td>Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/</td>
        <td>false</td>
      </tr><tr>
        <td><b>resourceVersion</b></td>
        <td>string</td>
        <td>Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency</td>
        <td>false</td>
      </tr><tr>
        <td><b>uid</b></td>
        <td>string</td>
        <td>UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids</td>
        <td>false</td>
      </tr></tbody>
</table>

## app.m4d.ibm.com/v1alpha1

Resource Types:

- [Blueprint](#blueprint)

- [M4DApplication](#m4dapplication)

- [M4DBucket](#m4dbucket)

- [M4DModule](#m4dmodule)

- [Plotter](#plotter)




### Blueprint
<sup><sup>[↩ Parent](#app.m4d.ibm.com/v1alpha1 )</sup></sup>






Blueprint is the Schema for the blueprints API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>app.m4d.ibm.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>Blueprint</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#blueprintspec">spec</a></b></td>
        <td>object</td>
        <td>BlueprintSpec defines the desired state of Blueprint, which is the runtime environment which provides the Data Scientist's application with secure and governed access to the data requested in the M4DApplication. The blueprint uses an "argo like" syntax which indicates the components and the flow of data between them as steps TODO: Add an indication of the communication relationships between the components</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintstatus">status</a></b></td>
        <td>object</td>
        <td>BlueprintStatus defines the observed state of Blueprint This includes readiness, error message, and indicators forthe Kubernetes resources owned by the Blueprint for cleanup and status monitoring</td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec
<sup><sup>[↩ Parent](#blueprint)</sup></sup>



BlueprintSpec defines the desired state of Blueprint, which is the runtime environment which provides the Data Scientist's application with secure and governed access to the data requested in the M4DApplication. The blueprint uses an "argo like" syntax which indicates the components and the flow of data between them as steps TODO: Add an indication of the communication relationships between the components

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>entrypoint</b></td>
        <td>string</td>
        <td></td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflow">flow</a></b></td>
        <td>object</td>
        <td>DataFlow indicates the flow of the data between the components Currently we assume this is linear and thus use steps, but other more complex graphs could be defined as per how it is done in argo workflow</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#blueprintspecselector">selector</a></b></td>
        <td>object</td>
        <td>Selector enables to connect the resource to the application Should match the selector of the owner - M4DApplication CRD.</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#blueprintspectemplatesindex">templates</a></b></td>
        <td>[]object</td>
        <td></td>
        <td>true</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow
<sup><sup>[↩ Parent](#blueprintspec)</sup></sup>



DataFlow indicates the flow of the data between the components Currently we assume this is linear and thus use steps, but other more complex graphs could be defined as per how it is done in argo workflow

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindex">steps</a></b></td>
        <td>[]object</td>
        <td></td>
        <td>true</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index]
<sup><sup>[↩ Parent](#blueprintspecflow)</sup></sup>



FlowStep is one step indicates an instance of a module in the blueprint, It includes the name of the module template (spec) and the parameters received by the component instance that is initiated by the orchestrator.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecflowstepsindexarguments">arguments</a></b></td>
        <td>object</td>
        <td>Arguments are the input parameters for a specific instance of a module.</td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>Name is the name of the instance of the module. For example, if the application is named "notebook" and an implicitcopy module is deemed necessary.  The FlowStep name would be notebook-implicitcopy.</td>
        <td>true</td>
      </tr><tr>
        <td><b>template</b></td>
        <td>string</td>
        <td>Template is the name of the specification in the Blueprint describing how to instantiate a component indicated by the module.  It is the name of a M4DModule CRD. For example: implicit-copy-db2wh-to-s3-latest</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments
<sup><sup>[↩ Parent](#blueprintspecflowstepsindex)</sup></sup>



Arguments are the input parameters for a specific instance of a module.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentscopy">copy</a></b></td>
        <td>object</td>
        <td>CopyArgs are parameters specific to modules that copy data from one data store to another.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentsreadindex">read</a></b></td>
        <td>[]object</td>
        <td>ReadArgs are parameters that are specific to modules that enable an application to read data</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentswriteindex">write</a></b></td>
        <td>[]object</td>
        <td>WriteArgs are parameters that are specific to modules that enable an application to write data</td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.copy
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexarguments)</sup></sup>



CopyArgs are parameters specific to modules that copy data from one data store to another.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentscopytransformationsindex">transformations</a></b></td>
        <td>[]object</td>
        <td>Transformations are different types of processing that may be done to the data as it is copied.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentscopydestination">destination</a></b></td>
        <td>object</td>
        <td>Destination is the data store to which the data will be copied</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentscopysource">source</a></b></td>
        <td>object</td>
        <td>Source is the where the data currently resides</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.copy.transformations[index]
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentscopy)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>args</b></td>
        <td>map[string]string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>level</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.copy.destination
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentscopy)</sup></sup>



Destination is the data store to which the data will be copied

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentscopydestinationconnection">connection</a></b></td>
        <td>object</td>
        <td>Connection has the relevant details for accesing the data (url, table, ssl, etc.)</td>
        <td>true</td>
      </tr><tr>
        <td><b>credentialLocation</b></td>
        <td>string</td>
        <td>CredentialLocation is used to obtain the credentials from the credential management system - ex: vault</td>
        <td>true</td>
      </tr><tr>
        <td><b>format</b></td>
        <td>string</td>
        <td>Format represents data format (e.g. parquet) as received from catalog connectors</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.copy.destination.connection
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentscopydestination)</sup></sup>



Connection has the relevant details for accesing the data (url, table, ssl, etc.)

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentscopydestinationconnectiondb2">db2</a></b></td>
        <td>object</td>
        <td>oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentscopydestinationconnectionkafka">kafka</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentscopydestinationconnections3">s3</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.copy.destination.connection.db2
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentscopydestinationconnection)</sup></sup>



oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>database</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>port</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>table</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.copy.destination.connection.kafka
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentscopydestinationconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bootstrap_servers</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>key_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>sasl_mechanism</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>schema_registry</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>security_protocol</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore_password</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>topic_name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>value_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.copy.destination.connection.s3
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentscopydestinationconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bucket</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>object_key</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>region</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.copy.source
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentscopy)</sup></sup>



Source is the where the data currently resides

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentscopysourceconnection">connection</a></b></td>
        <td>object</td>
        <td>Connection has the relevant details for accesing the data (url, table, ssl, etc.)</td>
        <td>true</td>
      </tr><tr>
        <td><b>credentialLocation</b></td>
        <td>string</td>
        <td>CredentialLocation is used to obtain the credentials from the credential management system - ex: vault</td>
        <td>true</td>
      </tr><tr>
        <td><b>format</b></td>
        <td>string</td>
        <td>Format represents data format (e.g. parquet) as received from catalog connectors</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.copy.source.connection
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentscopysource)</sup></sup>



Connection has the relevant details for accesing the data (url, table, ssl, etc.)

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentscopysourceconnectiondb2">db2</a></b></td>
        <td>object</td>
        <td>oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentscopysourceconnectionkafka">kafka</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentscopysourceconnections3">s3</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.copy.source.connection.db2
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentscopysourceconnection)</sup></sup>



oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>database</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>port</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>table</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.copy.source.connection.kafka
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentscopysourceconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bootstrap_servers</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>key_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>sasl_mechanism</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>schema_registry</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>security_protocol</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore_password</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>topic_name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>value_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.copy.source.connection.s3
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentscopysourceconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bucket</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>object_key</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>region</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.read[index]
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexarguments)</sup></sup>



ReadModuleArgs define the input parameters for modules that read data from location A

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentsreadindextransformationsindex">transformations</a></b></td>
        <td>[]object</td>
        <td>Transformations are different types of processing that may be done to the data</td>
        <td>false</td>
      </tr><tr>
        <td><b>assetID</b></td>
        <td>string</td>
        <td>AssetID identifies the asset to be used for accessing the data when it is ready It is copied from the M4DApplication resource</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentsreadindexsource">source</a></b></td>
        <td>object</td>
        <td>Source of the read path module</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.read[index].transformations[index]
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentsreadindex)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>args</b></td>
        <td>map[string]string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>level</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.read[index].source
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentsreadindex)</sup></sup>



Source of the read path module

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentsreadindexsourceconnection">connection</a></b></td>
        <td>object</td>
        <td>Connection has the relevant details for accesing the data (url, table, ssl, etc.)</td>
        <td>true</td>
      </tr><tr>
        <td><b>credentialLocation</b></td>
        <td>string</td>
        <td>CredentialLocation is used to obtain the credentials from the credential management system - ex: vault</td>
        <td>true</td>
      </tr><tr>
        <td><b>format</b></td>
        <td>string</td>
        <td>Format represents data format (e.g. parquet) as received from catalog connectors</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.read[index].source.connection
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentsreadindexsource)</sup></sup>



Connection has the relevant details for accesing the data (url, table, ssl, etc.)

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentsreadindexsourceconnectiondb2">db2</a></b></td>
        <td>object</td>
        <td>oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentsreadindexsourceconnectionkafka">kafka</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentsreadindexsourceconnections3">s3</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.read[index].source.connection.db2
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentsreadindexsourceconnection)</sup></sup>



oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>database</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>port</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>table</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.read[index].source.connection.kafka
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentsreadindexsourceconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bootstrap_servers</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>key_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>sasl_mechanism</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>schema_registry</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>security_protocol</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore_password</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>topic_name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>value_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.read[index].source.connection.s3
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentsreadindexsourceconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bucket</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>object_key</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>region</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.write[index]
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexarguments)</sup></sup>



WriteModuleArgs define the input parameters for modules that write data to location B

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentswriteindextransformationsindex">transformations</a></b></td>
        <td>[]object</td>
        <td>Transformations are different types of processing that may be done to the data as it is written.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentswriteindexdestination">destination</a></b></td>
        <td>object</td>
        <td>Destination is the data store to which the data will be written</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.write[index].transformations[index]
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentswriteindex)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>args</b></td>
        <td>map[string]string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>level</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.write[index].destination
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentswriteindex)</sup></sup>



Destination is the data store to which the data will be written

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentswriteindexdestinationconnection">connection</a></b></td>
        <td>object</td>
        <td>Connection has the relevant details for accesing the data (url, table, ssl, etc.)</td>
        <td>true</td>
      </tr><tr>
        <td><b>credentialLocation</b></td>
        <td>string</td>
        <td>CredentialLocation is used to obtain the credentials from the credential management system - ex: vault</td>
        <td>true</td>
      </tr><tr>
        <td><b>format</b></td>
        <td>string</td>
        <td>Format represents data format (e.g. parquet) as received from catalog connectors</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.write[index].destination.connection
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentswriteindexdestination)</sup></sup>



Connection has the relevant details for accesing the data (url, table, ssl, etc.)

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentswriteindexdestinationconnectiondb2">db2</a></b></td>
        <td>object</td>
        <td>oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentswriteindexdestinationconnectionkafka">kafka</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintspecflowstepsindexargumentswriteindexdestinationconnections3">s3</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.write[index].destination.connection.db2
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentswriteindexdestinationconnection)</sup></sup>



oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>database</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>port</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>table</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.write[index].destination.connection.kafka
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentswriteindexdestinationconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bootstrap_servers</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>key_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>sasl_mechanism</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>schema_registry</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>security_protocol</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore_password</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>topic_name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>value_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.flow.steps[index].arguments.write[index].destination.connection.s3
<sup><sup>[↩ Parent](#blueprintspecflowstepsindexargumentswriteindexdestinationconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bucket</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>object_key</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>region</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.selector
<sup><sup>[↩ Parent](#blueprintspec)</sup></sup>



Selector enables to connect the resource to the application Should match the selector of the owner - M4DApplication CRD.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#blueprintspecselectormatchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>matchExpressions is a list of label selector requirements. The requirements are ANDed.</td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabels</b></td>
        <td>map[string]string</td>
        <td>matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.</td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.spec.selector.matchExpressions[index]
<sup><sup>[↩ Parent](#blueprintspecselector)</sup></sup>



A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.</td>
        <td>false</td>
      </tr><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>key is the label key that the selector applies to.</td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Blueprint.spec.templates[index]
<sup><sup>[↩ Parent](#blueprintspec)</sup></sup>



ComponentTemplate is a copy of a M4DModule Custom Resource.  It contains the information necessary to instantiate a component in a FlowStep, which provides the functionality described by the module.  There are 3 different module types.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>kind</b></td>
        <td>string</td>
        <td>Kind of k8s resource</td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>Name of the template</td>
        <td>true</td>
      </tr><tr>
        <td><b>resources</b></td>
        <td>[]string</td>
        <td>Resources contains the location of the helm chart with info detailing how to deploy</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Blueprint.status
<sup><sup>[↩ Parent](#blueprint)</sup></sup>



BlueprintStatus defines the observed state of Blueprint This includes readiness, error message, and indicators forthe Kubernetes resources owned by the Blueprint for cleanup and status monitoring

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>observedGeneration</b></td>
        <td>integer</td>
        <td>ObservedGeneration is taken from the Blueprint metadata.  This is used to determine during reconcile whether reconcile was called because the desired state changed, or whether status of the allocated resources should be checked.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#blueprintstatusobservedstate">observedState</a></b></td>
        <td>object</td>
        <td>ObservedState includes information to be reported back to the M4DApplication resource It includes readiness and error indications, as well as user instructions</td>
        <td>false</td>
      </tr></tbody>
</table>


#### Blueprint.status.observedState
<sup><sup>[↩ Parent](#blueprintstatus)</sup></sup>



ObservedState includes information to be reported back to the M4DApplication resource It includes readiness and error indications, as well as user instructions

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>dataAccessInstructions</b></td>
        <td>string</td>
        <td>DataAccessInstructions indicate how the data user or his application may access the data. Instructions are available upon successful orchestration.</td>
        <td>false</td>
      </tr><tr>
        <td><b>error</b></td>
        <td>string</td>
        <td>Error indicates that there has been an error to orchestrate the modules and provides the error message</td>
        <td>false</td>
      </tr><tr>
        <td><b>ready</b></td>
        <td>boolean</td>
        <td>Ready represents that the modules have been orchestrated successfully and the data is ready for usage</td>
        <td>false</td>
      </tr></tbody>
</table>

### M4DApplication
<sup><sup>[↩ Parent](#app.m4d.ibm.com/v1alpha1 )</sup></sup>






M4DApplication provides information about the application being used by a Data Scientist, the nature of the processing, and the data sets that the Data Scientist has chosen for processing by the application. The M4DApplication controller (aka pilot) obtains instructions regarding any governance related changes that must be performed on the data, identifies the modules capable of performing such changes, and finally generates the Blueprint which defines the secure runtime environment and all the components in it.  This runtime environment provides the Data Scientist's application with access to the data requested in a secure manner and without having to provide any credentials for the data sets.  The credentials are obtained automatically by the manager from an external credential management system, which may or may not be part of a data catalog.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>app.m4d.ibm.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>M4DApplication</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#m4dapplicationspec">spec</a></b></td>
        <td>object</td>
        <td>M4DApplicationSpec defines the desired state of M4DApplication.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#m4dapplicationstatus">status</a></b></td>
        <td>object</td>
        <td>M4DApplicationStatus defines the observed state of M4DApplication.</td>
        <td>false</td>
      </tr></tbody>
</table>


#### M4DApplication.spec
<sup><sup>[↩ Parent](#m4dapplication)</sup></sup>



M4DApplicationSpec defines the desired state of M4DApplication.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#m4dapplicationspecappinfo">appInfo</a></b></td>
        <td>object</td>
        <td>AppInfo contains information describing the reasons and geography of the processing that will be done by the Data Scientist's application.</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#m4dapplicationspecdataindex">data</a></b></td>
        <td>[]object</td>
        <td>Data contains the identifiers of the data to be used by the Data Scientist's application, and the protocol used to access it and the format expected.</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#m4dapplicationspecselector">selector</a></b></td>
        <td>object</td>
        <td>Selector enables to connect the resource to the application Application labels should match the labels in the selector.</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DApplication.spec.appInfo
<sup><sup>[↩ Parent](#m4dapplicationspec)</sup></sup>



AppInfo contains information describing the reasons and geography of the processing that will be done by the Data Scientist's application.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>processingGeography</b></td>
        <td>string</td>
        <td>ProcessingGeography indicates the state or country or union in which the data processing will take place. This should be the same as the location of the cluster in which the manager is deployed.</td>
        <td>false</td>
      </tr><tr>
        <td><b>purpose</b></td>
        <td>string</td>
        <td>Purpose indicates the reason for the processing and the use of the data by the Data Scientist's application.</td>
        <td>false</td>
      </tr><tr>
        <td><b>role</b></td>
        <td>string</td>
        <td>Role indicates the position held or role filled by the Data Scientist as it relates to the processing of the data he has chosen.</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DApplication.spec.data[index]
<sup><sup>[↩ Parent](#m4dapplicationspec)</sup></sup>



DataContext indicates data set chosen by the Data Scientist to be used by his application, and includes information about the data format and technologies used by the application to access the data.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>dataSetID</b></td>
        <td>string</td>
        <td>DataSetID is a unique identifier of the dataset chosen from the data catalog for processing by the data user application.</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#m4dapplicationspecdataindexifdetails">ifDetails</a></b></td>
        <td>object</td>
        <td>IFdetails indicates the protocol and format expected by the data by the Data Scientist's application</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DApplication.spec.data[index].ifDetails
<sup><sup>[↩ Parent](#m4dapplicationspecdataindex)</sup></sup>



IFdetails indicates the protocol and format expected by the data by the Data Scientist's application

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>dataformat</b></td>
        <td>enum</td>
        <td>DataFormatType defines data format type [parquet table csv json avro binary arrow]</td>
        <td>false</td>
      </tr><tr>
        <td><b>protocol</b></td>
        <td>enum</td>
        <td>IFProtocol defines interface protocol for data transactions [s3 kafka jdbc-db2 m4d-arrow-flight]</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DApplication.spec.selector
<sup><sup>[↩ Parent](#m4dapplicationspec)</sup></sup>



Selector enables to connect the resource to the application Application labels should match the labels in the selector.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#m4dapplicationspecselectormatchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>matchExpressions is a list of label selector requirements. The requirements are ANDed.</td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabels</b></td>
        <td>map[string]string</td>
        <td>matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.</td>
        <td>false</td>
      </tr></tbody>
</table>


#### M4DApplication.spec.selector.matchExpressions[index]
<sup><sup>[↩ Parent](#m4dapplicationspecselector)</sup></sup>



A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.</td>
        <td>false</td>
      </tr><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>key is the label key that the selector applies to.</td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DApplication.status
<sup><sup>[↩ Parent](#m4dapplication)</sup></sup>



M4DApplicationStatus defines the observed state of M4DApplication.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#m4dapplicationstatusconditionsindex">conditions</a></b></td>
        <td>[]object</td>
        <td>Conditions represent the possible error and failure conditions</td>
        <td>false</td>
      </tr><tr>
        <td><b>dataAccessInstructions</b></td>
        <td>string</td>
        <td>DataAccessInstructions indicate how the data user or his application may access the data. Instructions are available upon successful orchestration.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#m4dapplicationstatusgenerated">generated</a></b></td>
        <td>object</td>
        <td>Generated resource identifier</td>
        <td>false</td>
      </tr><tr>
        <td><b>observedGeneration</b></td>
        <td>integer</td>
        <td>ObservedGeneration is taken from the M4DApplication metadata.  This is used to determine during reconcile whether reconcile was called because the desired state changed, or whether the Blueprint status changed.</td>
        <td>false</td>
      </tr><tr>
        <td><b>ready</b></td>
        <td>boolean</td>
        <td>Ready is true if a blueprint has been successfully orchestrated</td>
        <td>false</td>
      </tr></tbody>
</table>


#### M4DApplication.status.conditions[index]
<sup><sup>[↩ Parent](#m4dapplicationstatus)</sup></sup>



Condition describes the state of a M4DApplication at a certain point.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>message</b></td>
        <td>string</td>
        <td>Message contains the details of the current condition</td>
        <td>false</td>
      </tr><tr>
        <td><b>status</b></td>
        <td>string</td>
        <td>Status of the condition: true or false</td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>Type of the condition</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DApplication.status.generated
<sup><sup>[↩ Parent](#m4dapplicationstatus)</sup></sup>



Generated resource identifier

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>kind</b></td>
        <td>string</td>
        <td>Kind of the resource (Blueprint, Plotter)</td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>Name of the resource</td>
        <td>true</td>
      </tr><tr>
        <td><b>namespace</b></td>
        <td>string</td>
        <td>Namespace of the resource</td>
        <td>true</td>
      </tr></tbody>
</table>

### M4DBucket
<sup><sup>[↩ Parent](#app.m4d.ibm.com/v1alpha1 )</sup></sup>






M4DBucket defines a storage asset used for implicit copy destination It contains endpoint, bucket name, asset name and vault path where credentials are stored Owner of the asset is responsible to store the credentials in vault

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>app.m4d.ibm.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>M4DBucket</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#m4dbucketspec">spec</a></b></td>
        <td>object</td>
        <td>M4DBucketSpec defines the desired state of M4DBucket</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#m4dbucketstatus">status</a></b></td>
        <td>object</td>
        <td>M4DBucketStatus defines the observed state of M4DBucket</td>
        <td>false</td>
      </tr></tbody>
</table>


#### M4DBucket.spec
<sup><sup>[↩ Parent](#m4dbucket)</sup></sup>



M4DBucketSpec defines the desired state of M4DBucket

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td>Endpoint</td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>Provisioned bucket name</td>
        <td>true</td>
      </tr><tr>
        <td><b>vaultPath</b></td>
        <td>string</td>
        <td>Path where the credentials are stored</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DBucket.status
<sup><sup>[↩ Parent](#m4dbucket)</sup></sup>



M4DBucketStatus defines the observed state of M4DBucket

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>assetPrefixPerDataset</b></td>
        <td>map[string]string</td>
        <td>Each data asset for which the bucket is provisioned is mapped to the destination data asset (prefix) This is used for sharing a single bucket by multiple applications for the same data asset The data asset is identified by a string combined from dataset and catalog ids</td>
        <td>false</td>
      </tr><tr>
        <td><b>owners</b></td>
        <td>[]string</td>
        <td>Owner list: each resource is identified by namespace/name</td>
        <td>false</td>
      </tr></tbody>
</table>

### M4DModule
<sup><sup>[↩ Parent](#app.m4d.ibm.com/v1alpha1 )</sup></sup>






M4DModule is a description of an injectable component. the parameters it requires, as well as the specification of how to instantiate such a component. It is used as metadata only.  There is no status nor reconciliation.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>app.m4d.ibm.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>M4DModule</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#m4dmodulespec">spec</a></b></td>
        <td>object</td>
        <td>M4DModuleSpec contains the info common to all modules, which are one of the components that process, load, write, audit, monitor the data used by the data scientist's application.</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DModule.spec
<sup><sup>[↩ Parent](#m4dmodule)</sup></sup>



M4DModuleSpec contains the info common to all modules, which are one of the components that process, load, write, audit, monitor the data used by the data scientist's application.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#m4dmodulespecdependenciesindex">dependencies</a></b></td>
        <td>[]object</td>
        <td>Other components that must be installed in order for this module to work</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#m4dmodulespecstatusindicatorsindex">statusIndicators</a></b></td>
        <td>[]object</td>
        <td>StatusIndicators allow to check status of a non-standard resource that can not be computed by helm/kstatus</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#m4dmodulespeccapabilities">capabilities</a></b></td>
        <td>object</td>
        <td>Capabilities declares what this module knows how to do and the types of data it knows how to handle</td>
        <td>true</td>
      </tr><tr>
        <td><b>chart</b></td>
        <td>string</td>
        <td>Reference to a Helm chart that allows deployment of the resources required for this module</td>
        <td>true</td>
      </tr><tr>
        <td><b>flows</b></td>
        <td>[]enum</td>
        <td>Flows is a list of the types of capabilities supported by the module - copy, read, write</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DModule.spec.dependencies[index]
<sup><sup>[↩ Parent](#m4dmodulespec)</sup></sup>



Dependency details another component on which this module relies - i.e. a pre-requisit

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>Name is the name of the dependent component</td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>enum</td>
        <td>Type provides information used in determining how to instantiate the component [module connector feature]</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DModule.spec.statusIndicators[index]
<sup><sup>[↩ Parent](#m4dmodulespec)</sup></sup>



ResourceStatusIndicator is used to determine the status of an orchestrated resource

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>errorMessage</b></td>
        <td>string</td>
        <td>ErrorMessage specifies the resource field to check for an error, e.g. status.errorMsg</td>
        <td>false</td>
      </tr><tr>
        <td><b>failureCondition</b></td>
        <td>string</td>
        <td>FailureCondition specifies a condition that indicates the resource failure It uses kubernetes label selection syntax (https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)</td>
        <td>false</td>
      </tr><tr>
        <td><b>kind</b></td>
        <td>string</td>
        <td>Kind provides information about the resource kind</td>
        <td>true</td>
      </tr><tr>
        <td><b>successCondition</b></td>
        <td>string</td>
        <td>SuccessCondition specifies a condition that indicates that the resource is ready It uses kubernetes label selection syntax (https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DModule.spec.capabilities
<sup><sup>[↩ Parent](#m4dmodulespec)</sup></sup>



Capabilities declares what this module knows how to do and the types of data it knows how to handle

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#m4dmodulespeccapabilitiesactionsindex">actions</a></b></td>
        <td>[]object</td>
        <td>Actions are the data transformations that the module supports</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#m4dmodulespeccapabilitiesapi">api</a></b></td>
        <td>object</td>
        <td>API indicates to the application how to access/write the data</td>
        <td>false</td>
      </tr><tr>
        <td><b>credentials-managed-by</b></td>
        <td>enum</td>
        <td>Type provides information used in determining how to instantiate the component [secret-provider automatic]</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#m4dmodulespeccapabilitiessupportedinterfacesindex">supportedInterfaces</a></b></td>
        <td>[]object</td>
        <td>Copy should have one or more instances in the list, and its content should have source and sink Read should have one or more instances in the list, each with source populated Write should have one or more instances in the list, each with sink populated TODO - In the future if we have a module type that doesn't interface directly with data then this list could be empty</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DModule.spec.capabilities.actions[index]
<sup><sup>[↩ Parent](#m4dmodulespeccapabilities)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>args</b></td>
        <td>map[string]string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>level</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### M4DModule.spec.capabilities.api
<sup><sup>[↩ Parent](#m4dmodulespeccapabilities)</sup></sup>



API indicates to the application how to access/write the data

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>dataformat</b></td>
        <td>enum</td>
        <td>DataFormatType defines data format type [parquet table csv json avro binary arrow]</td>
        <td>false</td>
      </tr><tr>
        <td><b>protocol</b></td>
        <td>enum</td>
        <td>IFProtocol defines interface protocol for data transactions [s3 kafka jdbc-db2 m4d-arrow-flight]</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DModule.spec.capabilities.supportedInterfaces[index]
<sup><sup>[↩ Parent](#m4dmodulespeccapabilities)</sup></sup>



ModuleInOut specifies the protocol and format of the data input and output by the module - if any

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#m4dmodulespeccapabilitiessupportedinterfacesindexsink">sink</a></b></td>
        <td>object</td>
        <td>Sink specifies the output data protocol and format</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#m4dmodulespeccapabilitiessupportedinterfacesindexsource">source</a></b></td>
        <td>object</td>
        <td>Source specifies the input data protocol and format</td>
        <td>false</td>
      </tr><tr>
        <td><b>flow</b></td>
        <td>enum</td>
        <td>Flow for which this interface is supported [copy read write]</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DModule.spec.capabilities.supportedInterfaces[index].sink
<sup><sup>[↩ Parent](#m4dmodulespeccapabilitiessupportedinterfacesindex)</sup></sup>



Sink specifies the output data protocol and format

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>dataformat</b></td>
        <td>enum</td>
        <td>DataFormatType defines data format type [parquet table csv json avro binary arrow]</td>
        <td>false</td>
      </tr><tr>
        <td><b>protocol</b></td>
        <td>enum</td>
        <td>IFProtocol defines interface protocol for data transactions [s3 kafka jdbc-db2 m4d-arrow-flight]</td>
        <td>true</td>
      </tr></tbody>
</table>


#### M4DModule.spec.capabilities.supportedInterfaces[index].source
<sup><sup>[↩ Parent](#m4dmodulespeccapabilitiessupportedinterfacesindex)</sup></sup>



Source specifies the input data protocol and format

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>dataformat</b></td>
        <td>enum</td>
        <td>DataFormatType defines data format type [parquet table csv json avro binary arrow]</td>
        <td>false</td>
      </tr><tr>
        <td><b>protocol</b></td>
        <td>enum</td>
        <td>IFProtocol defines interface protocol for data transactions [s3 kafka jdbc-db2 m4d-arrow-flight]</td>
        <td>true</td>
      </tr></tbody>
</table>

### Plotter
<sup><sup>[↩ Parent](#app.m4d.ibm.com/v1alpha1 )</sup></sup>






Plotter is the Schema for the plotters API

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>app.m4d.ibm.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>Plotter</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#plotterspec">spec</a></b></td>
        <td>object</td>
        <td>PlotterSpec defines the desired state of Plotter, which is applied in a multi-clustered environment. Plotter installs the runtime environment (as blueprints running on remote clusters) which provides the Data Scientist's application with secure and governed access to the data requested in the M4DApplication.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterstatus">status</a></b></td>
        <td>object</td>
        <td>PlotterStatus defines the observed state of Plotter This includes readiness, error message, and indicators received from blueprint resources owned by the Plotter for cleanup and status monitoring</td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec
<sup><sup>[↩ Parent](#plotter)</sup></sup>



PlotterSpec defines the desired state of Plotter, which is applied in a multi-clustered environment. Plotter installs the runtime environment (as blueprints running on remote clusters) which provides the Data Scientist's application with secure and governed access to the data requested in the M4DApplication.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskey">blueprints</a></b></td>
        <td>map[string]object</td>
        <td>Blueprints structure represents remote blueprints mapped by the identifier of a cluster in which they will be running</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#plotterspecselector">selector</a></b></td>
        <td>object</td>
        <td>Selector enables to connect the resource to the application Should match the selector of the owner - M4DApplication CRD.</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key]
<sup><sup>[↩ Parent](#plotterspec)</sup></sup>



BlueprintSpec defines the desired state of Blueprint, which is the runtime environment which provides the Data Scientist's application with secure and governed access to the data requested in the M4DApplication. The blueprint uses an "argo like" syntax which indicates the components and the flow of data between them as steps TODO: Add an indication of the communication relationships between the components

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>entrypoint</b></td>
        <td>string</td>
        <td></td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflow">flow</a></b></td>
        <td>object</td>
        <td>DataFlow indicates the flow of the data between the components Currently we assume this is linear and thus use steps, but other more complex graphs could be defined as per how it is done in argo workflow</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyselector">selector</a></b></td>
        <td>object</td>
        <td>Selector enables to connect the resource to the application Should match the selector of the owner - M4DApplication CRD.</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeytemplatesindex">templates</a></b></td>
        <td>[]object</td>
        <td></td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow
<sup><sup>[↩ Parent](#plotterspecblueprintskey)</sup></sup>



DataFlow indicates the flow of the data between the components Currently we assume this is linear and thus use steps, but other more complex graphs could be defined as per how it is done in argo workflow

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindex">steps</a></b></td>
        <td>[]object</td>
        <td></td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index]
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflow)</sup></sup>



FlowStep is one step indicates an instance of a module in the blueprint, It includes the name of the module template (spec) and the parameters received by the component instance that is initiated by the orchestrator.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexarguments">arguments</a></b></td>
        <td>object</td>
        <td>Arguments are the input parameters for a specific instance of a module.</td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>Name is the name of the instance of the module. For example, if the application is named "notebook" and an implicitcopy module is deemed necessary.  The FlowStep name would be notebook-implicitcopy.</td>
        <td>true</td>
      </tr><tr>
        <td><b>template</b></td>
        <td>string</td>
        <td>Template is the name of the specification in the Blueprint describing how to instantiate a component indicated by the module.  It is the name of a M4DModule CRD. For example: implicit-copy-db2wh-to-s3-latest</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindex)</sup></sup>



Arguments are the input parameters for a specific instance of a module.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentscopy">copy</a></b></td>
        <td>object</td>
        <td>CopyArgs are parameters specific to modules that copy data from one data store to another.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentsreadindex">read</a></b></td>
        <td>[]object</td>
        <td>ReadArgs are parameters that are specific to modules that enable an application to read data</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentswriteindex">write</a></b></td>
        <td>[]object</td>
        <td>WriteArgs are parameters that are specific to modules that enable an application to write data</td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.copy
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexarguments)</sup></sup>



CopyArgs are parameters specific to modules that copy data from one data store to another.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentscopytransformationsindex">transformations</a></b></td>
        <td>[]object</td>
        <td>Transformations are different types of processing that may be done to the data as it is copied.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentscopydestination">destination</a></b></td>
        <td>object</td>
        <td>Destination is the data store to which the data will be copied</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentscopysource">source</a></b></td>
        <td>object</td>
        <td>Source is the where the data currently resides</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.copy.transformations[index]
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentscopy)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>args</b></td>
        <td>map[string]string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>level</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.copy.destination
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentscopy)</sup></sup>



Destination is the data store to which the data will be copied

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentscopydestinationconnection">connection</a></b></td>
        <td>object</td>
        <td>Connection has the relevant details for accesing the data (url, table, ssl, etc.)</td>
        <td>true</td>
      </tr><tr>
        <td><b>credentialLocation</b></td>
        <td>string</td>
        <td>CredentialLocation is used to obtain the credentials from the credential management system - ex: vault</td>
        <td>true</td>
      </tr><tr>
        <td><b>format</b></td>
        <td>string</td>
        <td>Format represents data format (e.g. parquet) as received from catalog connectors</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.copy.destination.connection
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentscopydestination)</sup></sup>



Connection has the relevant details for accesing the data (url, table, ssl, etc.)

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentscopydestinationconnectiondb2">db2</a></b></td>
        <td>object</td>
        <td>oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentscopydestinationconnectionkafka">kafka</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentscopydestinationconnections3">s3</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.copy.destination.connection.db2
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentscopydestinationconnection)</sup></sup>



oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>database</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>port</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>table</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.copy.destination.connection.kafka
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentscopydestinationconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bootstrap_servers</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>key_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>sasl_mechanism</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>schema_registry</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>security_protocol</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore_password</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>topic_name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>value_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.copy.destination.connection.s3
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentscopydestinationconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bucket</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>object_key</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>region</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.copy.source
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentscopy)</sup></sup>



Source is the where the data currently resides

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentscopysourceconnection">connection</a></b></td>
        <td>object</td>
        <td>Connection has the relevant details for accesing the data (url, table, ssl, etc.)</td>
        <td>true</td>
      </tr><tr>
        <td><b>credentialLocation</b></td>
        <td>string</td>
        <td>CredentialLocation is used to obtain the credentials from the credential management system - ex: vault</td>
        <td>true</td>
      </tr><tr>
        <td><b>format</b></td>
        <td>string</td>
        <td>Format represents data format (e.g. parquet) as received from catalog connectors</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.copy.source.connection
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentscopysource)</sup></sup>



Connection has the relevant details for accesing the data (url, table, ssl, etc.)

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentscopysourceconnectiondb2">db2</a></b></td>
        <td>object</td>
        <td>oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentscopysourceconnectionkafka">kafka</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentscopysourceconnections3">s3</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.copy.source.connection.db2
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentscopysourceconnection)</sup></sup>



oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>database</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>port</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>table</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.copy.source.connection.kafka
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentscopysourceconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bootstrap_servers</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>key_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>sasl_mechanism</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>schema_registry</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>security_protocol</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore_password</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>topic_name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>value_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.copy.source.connection.s3
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentscopysourceconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bucket</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>object_key</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>region</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.read[index]
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexarguments)</sup></sup>



ReadModuleArgs define the input parameters for modules that read data from location A

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentsreadindextransformationsindex">transformations</a></b></td>
        <td>[]object</td>
        <td>Transformations are different types of processing that may be done to the data</td>
        <td>false</td>
      </tr><tr>
        <td><b>assetID</b></td>
        <td>string</td>
        <td>AssetID identifies the asset to be used for accessing the data when it is ready It is copied from the M4DApplication resource</td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentsreadindexsource">source</a></b></td>
        <td>object</td>
        <td>Source of the read path module</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.read[index].transformations[index]
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentsreadindex)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>args</b></td>
        <td>map[string]string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>level</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.read[index].source
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentsreadindex)</sup></sup>



Source of the read path module

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentsreadindexsourceconnection">connection</a></b></td>
        <td>object</td>
        <td>Connection has the relevant details for accesing the data (url, table, ssl, etc.)</td>
        <td>true</td>
      </tr><tr>
        <td><b>credentialLocation</b></td>
        <td>string</td>
        <td>CredentialLocation is used to obtain the credentials from the credential management system - ex: vault</td>
        <td>true</td>
      </tr><tr>
        <td><b>format</b></td>
        <td>string</td>
        <td>Format represents data format (e.g. parquet) as received from catalog connectors</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.read[index].source.connection
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentsreadindexsource)</sup></sup>



Connection has the relevant details for accesing the data (url, table, ssl, etc.)

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentsreadindexsourceconnectiondb2">db2</a></b></td>
        <td>object</td>
        <td>oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentsreadindexsourceconnectionkafka">kafka</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentsreadindexsourceconnections3">s3</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.read[index].source.connection.db2
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentsreadindexsourceconnection)</sup></sup>



oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>database</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>port</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>table</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.read[index].source.connection.kafka
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentsreadindexsourceconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bootstrap_servers</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>key_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>sasl_mechanism</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>schema_registry</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>security_protocol</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore_password</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>topic_name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>value_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.read[index].source.connection.s3
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentsreadindexsourceconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bucket</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>object_key</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>region</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.write[index]
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexarguments)</sup></sup>



WriteModuleArgs define the input parameters for modules that write data to location B

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentswriteindextransformationsindex">transformations</a></b></td>
        <td>[]object</td>
        <td>Transformations are different types of processing that may be done to the data as it is written.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentswriteindexdestination">destination</a></b></td>
        <td>object</td>
        <td>Destination is the data store to which the data will be written</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.write[index].transformations[index]
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentswriteindex)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>args</b></td>
        <td>map[string]string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>id</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>level</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.write[index].destination
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentswriteindex)</sup></sup>



Destination is the data store to which the data will be written

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentswriteindexdestinationconnection">connection</a></b></td>
        <td>object</td>
        <td>Connection has the relevant details for accesing the data (url, table, ssl, etc.)</td>
        <td>true</td>
      </tr><tr>
        <td><b>credentialLocation</b></td>
        <td>string</td>
        <td>CredentialLocation is used to obtain the credentials from the credential management system - ex: vault</td>
        <td>true</td>
      </tr><tr>
        <td><b>format</b></td>
        <td>string</td>
        <td>Format represents data format (e.g. parquet) as received from catalog connectors</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.write[index].destination.connection
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentswriteindexdestination)</sup></sup>



Connection has the relevant details for accesing the data (url, table, ssl, etc.)

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentswriteindexdestinationconnectiondb2">db2</a></b></td>
        <td>object</td>
        <td>oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentswriteindexdestinationconnectionkafka">kafka</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterspecblueprintskeyflowstepsindexargumentswriteindexdestinationconnections3">s3</a></b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>integer</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.write[index].destination.connection.db2
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentswriteindexdestinationconnection)</sup></sup>



oneof location {   // should have been oneof but for technical rasons, a problem to translate it to JSON, we remove the oneof for now should have been local, db2, s3 without "location"  but had a problem to compile it in proto - collision with proto name DataLocationDb2

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>database</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>port</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>table</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.write[index].destination.connection.kafka
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentswriteindexdestinationconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bootstrap_servers</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>key_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>sasl_mechanism</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>schema_registry</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>security_protocol</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>ssl_truststore_password</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>topic_name</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>value_deserializer</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].flow.steps[index].arguments.write[index].destination.connection.s3
<sup><sup>[↩ Parent](#plotterspecblueprintskeyflowstepsindexargumentswriteindexdestinationconnection)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>bucket</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>endpoint</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>object_key</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>region</b></td>
        <td>string</td>
        <td></td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].selector
<sup><sup>[↩ Parent](#plotterspecblueprintskey)</sup></sup>



Selector enables to connect the resource to the application Should match the selector of the owner - M4DApplication CRD.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecblueprintskeyselectormatchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>matchExpressions is a list of label selector requirements. The requirements are ANDed.</td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabels</b></td>
        <td>map[string]string</td>
        <td>matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.</td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].selector.matchExpressions[index]
<sup><sup>[↩ Parent](#plotterspecblueprintskeyselector)</sup></sup>



A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.</td>
        <td>false</td>
      </tr><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>key is the label key that the selector applies to.</td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.spec.blueprints[key].templates[index]
<sup><sup>[↩ Parent](#plotterspecblueprintskey)</sup></sup>



ComponentTemplate is a copy of a M4DModule Custom Resource.  It contains the information necessary to instantiate a component in a FlowStep, which provides the functionality described by the module.  There are 3 different module types.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>kind</b></td>
        <td>string</td>
        <td>Kind of k8s resource</td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>Name of the template</td>
        <td>true</td>
      </tr><tr>
        <td><b>resources</b></td>
        <td>[]string</td>
        <td>Resources contains the location of the helm chart with info detailing how to deploy</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.spec.selector
<sup><sup>[↩ Parent](#plotterspec)</sup></sup>



Selector enables to connect the resource to the application Should match the selector of the owner - M4DApplication CRD.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterspecselectormatchexpressionsindex">matchExpressions</a></b></td>
        <td>[]object</td>
        <td>matchExpressions is a list of label selector requirements. The requirements are ANDed.</td>
        <td>false</td>
      </tr><tr>
        <td><b>matchLabels</b></td>
        <td>map[string]string</td>
        <td>matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.</td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.spec.selector.matchExpressions[index]
<sup><sup>[↩ Parent](#plotterspecselector)</sup></sup>



A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>values</b></td>
        <td>[]string</td>
        <td>values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.</td>
        <td>false</td>
      </tr><tr>
        <td><b>key</b></td>
        <td>string</td>
        <td>key is the label key that the selector applies to.</td>
        <td>true</td>
      </tr><tr>
        <td><b>operator</b></td>
        <td>string</td>
        <td>operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.</td>
        <td>true</td>
      </tr></tbody>
</table>


#### Plotter.status
<sup><sup>[↩ Parent](#plotter)</sup></sup>



PlotterStatus defines the observed state of Plotter This includes readiness, error message, and indicators received from blueprint resources owned by the Plotter for cleanup and status monitoring

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#plotterstatusblueprintskey">blueprints</a></b></td>
        <td>map[string]object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b>observedGeneration</b></td>
        <td>integer</td>
        <td>ObservedGeneration is taken from the Plotter metadata.  This is used to determine during reconcile whether reconcile was called because the desired state changed, or whether status of the allocated blueprints should be checked.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterstatusobservedstate">observedState</a></b></td>
        <td>object</td>
        <td>ObservedState includes information to be reported back to the M4DApplication resource It includes readiness and error indications, as well as user instructions</td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.status.blueprints[key]
<sup><sup>[↩ Parent](#plotterstatus)</sup></sup>





<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>metadata</b></td>
        <td>object</td>
        <td></td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterstatusblueprintskeystatus">status</a></b></td>
        <td>object</td>
        <td>BlueprintStatus defines the observed state of Blueprint This includes readiness, error message, and indicators forthe Kubernetes resources owned by the Blueprint for cleanup and status monitoring</td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.status.blueprints[key].status
<sup><sup>[↩ Parent](#plotterstatusblueprintskey)</sup></sup>



BlueprintStatus defines the observed state of Blueprint This includes readiness, error message, and indicators forthe Kubernetes resources owned by the Blueprint for cleanup and status monitoring

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>observedGeneration</b></td>
        <td>integer</td>
        <td>ObservedGeneration is taken from the Blueprint metadata.  This is used to determine during reconcile whether reconcile was called because the desired state changed, or whether status of the allocated resources should be checked.</td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#plotterstatusblueprintskeystatusobservedstate">observedState</a></b></td>
        <td>object</td>
        <td>ObservedState includes information to be reported back to the M4DApplication resource It includes readiness and error indications, as well as user instructions</td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.status.blueprints[key].status.observedState
<sup><sup>[↩ Parent](#plotterstatusblueprintskeystatus)</sup></sup>



ObservedState includes information to be reported back to the M4DApplication resource It includes readiness and error indications, as well as user instructions

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>dataAccessInstructions</b></td>
        <td>string</td>
        <td>DataAccessInstructions indicate how the data user or his application may access the data. Instructions are available upon successful orchestration.</td>
        <td>false</td>
      </tr><tr>
        <td><b>error</b></td>
        <td>string</td>
        <td>Error indicates that there has been an error to orchestrate the modules and provides the error message</td>
        <td>false</td>
      </tr><tr>
        <td><b>ready</b></td>
        <td>boolean</td>
        <td>Ready represents that the modules have been orchestrated successfully and the data is ready for usage</td>
        <td>false</td>
      </tr></tbody>
</table>


#### Plotter.status.observedState
<sup><sup>[↩ Parent](#plotterstatus)</sup></sup>



ObservedState includes information to be reported back to the M4DApplication resource It includes readiness and error indications, as well as user instructions

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>dataAccessInstructions</b></td>
        <td>string</td>
        <td>DataAccessInstructions indicate how the data user or his application may access the data. Instructions are available upon successful orchestration.</td>
        <td>false</td>
      </tr><tr>
        <td><b>error</b></td>
        <td>string</td>
        <td>Error indicates that there has been an error to orchestrate the modules and provides the error message</td>
        <td>false</td>
      </tr><tr>
        <td><b>ready</b></td>
        <td>boolean</td>
        <td>Ready represents that the modules have been orchestrated successfully and the data is ready for usage</td>
        <td>false</td>
      </tr></tbody>
</table>
