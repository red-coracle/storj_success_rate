package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

const (
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Red    = "\033[31m"
    Cyan   = "\033[36m"
    Reset  = "\033[0m"
)

type stat struct {
    success int
    warn int
    critical int
    rejected int
    canceled int
    failed int
}

type metrics struct {
    audit stat
    download stat
    upload stat
    repairdownload stat
    repairupload stat
    deletes stat
}

func (s stat) total() int {
    // Is there a way to iterate struct fields?
    return s.success + s.warn + s.critical + s.rejected + s.canceled + s.failed
}

func parse_audit(line string, results *stat) () {
    if strings.Contains(line, "downloaded") {
        results.success += 1
    } else if strings.Contains(line, "failed") {
        if !strings.Contains(line, "exist") {
            results.warn += 1
        } else {
            results.critical += 1
        }
    }
    return
}

func parse_store(line string, results *stat) () {
    if strings.Contains(line, "uploaded") || strings.Contains(line, "downloaded") {
        results.success += 1
    } else if strings.Contains(line, "rejected") {
        results.rejected += 1
    } else if strings.Contains(line, "canceled") {
        results.canceled += 1
    } else if strings.Contains(line, "failed") {
        results.failed += 1
    }
    return
}

func parse_delete(line string, results *stat) () {
    if strings.Contains(line, "deleted") || strings.Contains(line, "delete piece") {
        results.success += 1
    } else if strings.Contains(line, "delete failed") {
        results.failed += 1
    }
    return
}

func parse_file(path string) () {
    input, err := os.Open(path)
    if err != nil {
        return
    }
    defer input.Close()

    var results = metrics{}
    scanner := bufio.NewScanner(input)

    for scanner.Scan() {
        var candidate = scanner.Text()
        if strings.Contains(candidate, "GET_AUDIT") {
            parse_audit(candidate, &results.audit)
        } else if strings.Contains(candidate, "\"GET\"") {
            parse_store(candidate, &results.download)
        } else if strings.Contains(candidate, "\"PUT\"") {
            parse_store(candidate, &results.upload)
        } else if strings.Contains(candidate, "GET_REPAIR") {
            parse_store(candidate, &results.repairdownload)
        } else if strings.Contains(candidate, "PUT_REPAIR") {
            parse_store(candidate, &results.repairupload)
        } else {
            parse_delete(candidate, &results.deletes)
        }
    }

    fmt.Printf("%s========== AUDIT ==============%s\n", Cyan, Reset)
    fmt.Printf("%sCritically failed:     %d%s\n", Red, results.audit.critical, Reset)
    fmt.Printf("Critical Fail Rate:    %.3f%%\n", 100.0 * float64(results.audit.critical) / float64(results.audit.total()))
    fmt.Printf("%sRecoverable failed:    %d%s\n", Yellow, results.audit.warn, Reset)
    fmt.Printf("Recoverable Fail Rate: %.3f%%\n", 100.0 * float64(results.audit.warn) / float64(results.audit.total()))
    fmt.Printf("%sSuccessful:            %d%s\n", Green, results.audit.success, Reset)
    fmt.Printf("Success Rate:          %.3f%%\n", 100.0 * float64(results.audit.success) / float64(results.audit.total()))

    fmt.Printf("%s========== DOWNLOAD ===========%s\n", Cyan, Reset)
    fmt.Printf("%sFailed:                %d%s\n", Red, results.download.failed, Reset)
    fmt.Printf("Fail Rate:             %.3f%%\n", 100.0 * float64(results.download.failed) / float64(results.download.total()))
    fmt.Printf("%sCanceled:              %d%s\n", Yellow, results.download.canceled, Reset)
    fmt.Printf("Cancel Rate:           %.3f%%\n", 100.0 * float64(results.download.canceled) / float64(results.download.total()))
    fmt.Printf("%sSuccessful:            %d%s\n", Green, results.download.success, Reset)
    fmt.Printf("Success Rate:          %.3f%%\n", 100.0 * float64(results.download.success) / float64(results.download.total()))

    fmt.Printf("%s========== UPLOAD =============%s\n", Cyan, Reset)
    fmt.Printf("%sRejected:              %d%s\n", Yellow, results.upload.rejected, Reset)
    fmt.Printf("Acceptance Rate:       %.3f%%\n", 100.0 - 100.0 * float64(results.upload.rejected) / float64(results.upload.total()))
    fmt.Printf("%sFailed:                %d%s\n", Red, results.upload.failed, Reset)
    fmt.Printf("Fail Rate:             %.3f%%\n", 100.0 * float64(results.upload.failed) / float64(results.upload.total()))
    fmt.Printf("%sCanceled:              %d%s\n", Yellow, results.upload.canceled, Reset)
    fmt.Printf("Cancel Rate:           %.3f%%\n", 100.0 * float64(results.upload.canceled) / float64(results.upload.total()))
    fmt.Printf("%sSuccessful:            %d%s\n", Green, results.upload.success, Reset)
    fmt.Printf("Success Rate:          %.3f%%\n", 100.0 * float64(results.upload.success) / float64(results.upload.total()))

    fmt.Printf("%s========== REPAIR DOWNLOAD ====%s\n", Cyan, Reset)
    fmt.Printf("%sFailed:                %d%s\n", Red, results.repairdownload.failed, Reset)
    fmt.Printf("Fail Rate:             %.3f%%\n", 100.0 * float64(results.repairdownload.failed) / float64(results.repairdownload.total()))
    fmt.Printf("%sCanceled:              %d%s\n", Yellow, results.repairdownload.canceled, Reset)
    fmt.Printf("Cancel Rate:           %.3f%%\n", 100.0 * float64(results.repairdownload.canceled) / float64(results.repairdownload.total()))
    fmt.Printf("%sSuccessful:            %d%s\n", Green, results.repairdownload.success, Reset)
    fmt.Printf("Success Rate:          %.3f%%\n", 100.0 * float64(results.repairdownload.success) / float64(results.repairdownload.total()))

    fmt.Printf("%s========== REPAIR UPLOAD ======%s\n", Cyan, Reset)
    fmt.Printf("%sFailed:                %d%s\n", Red, results.repairupload.failed, Reset)
    fmt.Printf("Fail Rate:             %.3f%%\n", 100.0 * float64(results.repairupload.failed) / float64(results.repairupload.total()))
    fmt.Printf("%sCanceled:              %d%s\n", Yellow, results.repairupload.canceled, Reset)
    fmt.Printf("Cancel Rate:           %.3f%%\n", 100.0 * float64(results.repairupload.canceled) / float64(results.repairupload.total()))
    fmt.Printf("%sSuccessful:            %d%s\n", Green, results.repairupload.success, Reset)
    fmt.Printf("Success Rate:          %.3f%%\n", 100.0 * float64(results.repairupload.success) / float64(results.repairupload.total()))

    fmt.Printf("%s========== DELETE =============%s\n", Cyan, Reset)
    fmt.Printf("%sFailed:                %d%s\n", Red, results.deletes.failed, Reset)
    fmt.Printf("Fail Rate:             %.3f%%\n", 100.0 * float64(results.deletes.failed) / float64(results.deletes.total()))
    fmt.Printf("%sSuccessful:            %d%s\n", Green, results.deletes.success, Reset)
    fmt.Printf("Success Rate:          %.3f%%\n", 100.0 * float64(results.deletes.success) / float64(results.deletes.total()))
}

func main() {
    if len(os.Args) < 2 {
        fmt.Printf("Usage: %s <path to log file>\n", os.Args[0])
    } else {
        parse_file(os.Args[1])
    }
}
