package main

const QUERY_BY_USER ="{\"selector\":{\"userId\":\"%s\"}, \"use_index\":[\"_design/scanDoc\", \"scan\"]}"

const QUERY_BY_SERIALID ="{\"selector\":{\"serialID\":\"%s\"}, \"use_index\":[\"_design/scanDoc\", \"scan\"]}"

const QUERY_BY_USER_AND_SERIALID ="{\"selector\":{\"userId\":\"%s\",\"serialID\":\"%s\"}, \"use_index\":[\"_design/scanDoc\", \"scan\"]}"

const QUERY_BY_DATES ="{\"selector\":{\"lastUpdatedat\":{\"$gte\":\"%s\",\"$lte\":\"%s\"}}, \"use_index\":[\"_design/scanDoc\", \"scan\"]}"

const QUERY_BY_SCANNED_RESULTS="{\"selector\":{\"scannedResult\":{\"$regex\":\"^%s\"}}}"

const QUERY_BY_DATES_WITH_SCANNED_RESULTS="{\"selector\":{\"lastUpdatedat\":{\"$gte\":\"%s\",\"$lte\":\"%s\"},\"scannedResult\":{\"$regex\":\"^%s\"}},\"use_index\":[\"_design/scanDoc\", \"scan\"]}"

const QUERY_BY_DATES_WITH_SCANNED_LOCATION="{\"selector\":{\"lastUpdatedat\":{\"$gte\":\"%s\",\"$lte\":\"%s\"},\"scanLocation\":\"%s\"},\"use_index\":[\"_design/scanDoc\", \"scan\"]}"


const QUERY_BY_DATES_WITH_SCANNED_RESULTS_AND_SCANNED_LOCATION="{\"selector\":{\"lastUpdatedat\":{\"$gte\":\"%s\",\"$lte\":\"%s\"},\"scannedResult\":{\"$regex\":\"^%s\"},\"scanLocation\":\"%s\"},\"use_index\":[\"_design/scanDoc\", \"scan\"]}"


