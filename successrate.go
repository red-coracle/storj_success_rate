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

func parse_audit(line string) (string) {
    if strings.Contains(line, "downloaded") {
        return "success"
    } else if strings.Contains(line, "failed") {
        if !strings.Contains(line, "exist") {
            return "warn"
        } else {
            return "critical"
        }
    }
    return "other"
}

func parse_store(line string) (string) {
    if strings.Contains(line, "uploaded") || strings.Contains(line, "downloaded") {
        return "success"
    } else if strings.Contains(line, "rejected") {
        return "rejected"
    } else if strings.Contains(line, "canceled") {
        return "canceled"
    } else if strings.Contains(line, "failed") {
        return "failed"
    }
    return "other"
}

func parse_delete(line string) (string) {
    if strings.Contains(line, "deleted") || strings.Contains(line, "delete piece") {
        return "success"
    } else if strings.Contains(line, "delete failed") {
        return "failed"
    }

    return "other"
}

func parse_file(path string) () {
    input, err := os.Open(path)
    if err != nil {
        return
    }
    defer input.Close()

    scanner := bufio.NewScanner(input)
    audit_results := make(map[string]int)
    download_results := make(map[string]int)
    upload_results := make(map[string]int)
    repair_upload_results := make(map[string]int)
    repair_download_results := make(map[string]int)
    delete_results := make(map[string]int)

    for scanner.Scan() {
        var candidate = scanner.Text()
        if strings.Contains(candidate, "GET_AUDIT") {
            audit_results[parse_audit(candidate)] += 1
            audit_results["total"] += 1
        } else if strings.Contains(candidate, "\"GET\"") {
            download_results[parse_store(candidate)] += 1
            download_results["total"] += 1
        } else if strings.Contains(candidate, "\"PUT\"") {
            upload_results[parse_store(candidate)] += 1
            upload_results["total"] += 1
        } else if strings.Contains(candidate, "GET_REPAIR") {
            repair_download_results[parse_store(candidate)] += 1
            repair_download_results["total"] += 1
        } else if strings.Contains(candidate, "PUT_REPAIR") {
            repair_upload_results[parse_store(candidate)] += 1
            repair_upload_results["total"] += 1
        } else {
            delete_results[parse_delete(candidate)] += 1
            delete_results["total"] += 1
        }
    }

    audit_results["total"] -= audit_results["other"]
    download_results["total"] -= download_results["other"]
    upload_results["total"] -= upload_results["other"]
    repair_download_results["total"] -= repair_download_results["other"]
    repair_upload_results["total"] -= repair_upload_results["other"]
    delete_results["total"] -= delete_results["other"]

    fmt.Printf("%s========== AUDIT ==============%s\n", Cyan, Reset)
    fmt.Printf("%sCritically failed:     %d%s\n", Red, audit_results["critical"], Reset)
    fmt.Printf("Critical Fail Rate:    %.3f%%\n", 100.0 * float64(audit_results["critical"]) / float64(audit_results["total"]))
    fmt.Printf("%sRecoverable failed:    %d%s\n", Yellow, audit_results["warn"], Reset)
    fmt.Printf("Recoverable Fail Rate: %.3f%%\n", 100.0 * float64(audit_results["warn"]) / float64(audit_results["total"]))
    fmt.Printf("%sSuccessful:            %d%s\n", Green, audit_results["success"], Reset)
    fmt.Printf("Success Rate:          %.3f%%\n", 100.0 * float64(audit_results["success"]) / float64(audit_results["total"]))

    fmt.Printf("%s========== DOWNLOAD ===========%s\n", Cyan, Reset)
    fmt.Printf("%sFailed:                %d%s\n", Red, download_results["failed"], Reset)
    fmt.Printf("Fail Rate:             %.3f%%\n", 100.0 * float64(download_results["failed"]) / float64(download_results["total"]))
    fmt.Printf("%sCanceled:              %d%s\n", Yellow, download_results["canceled"], Reset)
    fmt.Printf("Cancel Rate:           %.3f%%\n", 100.0 * float64(download_results["canceled"]) / float64(download_results["total"]))
    fmt.Printf("%sSuccessful:            %d%s\n", Green, download_results["success"], Reset)
    fmt.Printf("Success Rate:          %.3f%%\n", 100.0 * float64(download_results["success"]) / float64(download_results["total"]))

    fmt.Printf("%s========== UPLOAD =============%s\n", Cyan, Reset)
    fmt.Printf("%sRejected:              %d%s\n", Yellow, upload_results["rejected"], Reset)
    fmt.Printf("Acceptance Rate:       %.3f%%\n", 100.0 - 100.0 * float64(upload_results["rejected"]) / float64(upload_results["total"]))
    fmt.Printf("%sFailed:                %d%s\n", Red, upload_results["failed"], Reset)
    fmt.Printf("Fail Rate:             %.3f%%\n", 100.0 * float64(upload_results["failed"]) / float64(upload_results["total"]))
    fmt.Printf("%sCanceled:              %d%s\n", Yellow, upload_results["canceled"], Reset)
    fmt.Printf("Cancel Rate:           %.3f%%\n", 100.0 * float64(upload_results["canceled"]) / float64(upload_results["total"]))
    fmt.Printf("%sSuccessful:            %d%s\n", Green, upload_results["success"], Reset)
    fmt.Printf("Success Rate:          %.3f%%\n", 100.0 * float64(upload_results["success"]) / float64(upload_results["total"]))

    fmt.Printf("%s========== REPAIR DOWNLOAD ====%s\n", Cyan, Reset)
    fmt.Printf("%sFailed:                %d%s\n", Red, repair_download_results["failed"], Reset)
    fmt.Printf("Fail Rate:             %.3f%%\n", 100.0 * float64(repair_download_results["failed"]) / float64(repair_download_results["total"]))
    fmt.Printf("%sCanceled:              %d%s\n", Yellow, repair_download_results["canceled"], Reset)
    fmt.Printf("Cancel Rate:           %.3f%%\n", 100.0 * float64(repair_download_results["canceled"]) / float64(repair_download_results["total"]))
    fmt.Printf("%sSuccessful:            %d%s\n", Green, repair_download_results["success"], Reset)
    fmt.Printf("Success Rate:          %.3f%%\n", 100.0 * float64(repair_download_results["success"]) / float64(repair_download_results["total"]))

    fmt.Printf("%s========== REPAIR UPLOAD ======%s\n", Cyan, Reset)
    fmt.Printf("%sFailed:                %d%s\n", Red, repair_upload_results["failed"], Reset)
    fmt.Printf("Fail Rate:             %.3f%%\n", 100.0 * float64(repair_upload_results["failed"]) / float64(repair_upload_results["total"]))
    fmt.Printf("%sCanceled:              %d%s\n", Yellow, repair_upload_results["canceled"], Reset)
    fmt.Printf("Cancel Rate:           %.3f%%\n", 100.0 * float64(repair_upload_results["canceled"]) / float64(repair_upload_results["total"]))
    fmt.Printf("%sSuccessful:            %d%s\n", Green, repair_upload_results["success"], Reset)
    fmt.Printf("Success Rate:          %.3f%%\n", 100.0 * float64(repair_upload_results["success"]) / float64(repair_upload_results["total"]))

    fmt.Printf("%s========== DELETE =============%s\n", Cyan, Reset)
    fmt.Printf("%sFailed:                %d%s\n", Red, delete_results["failed"], Reset)
    fmt.Printf("Fail Rate:             %.3f%%\n", 100.0 * float64(delete_results["failed"]) / float64(delete_results["total"]))
    fmt.Printf("%sSuccessful:            %d%s\n", Green, delete_results["success"], Reset)
    fmt.Printf("Success Rate:          %.3f%%\n", 100.0 * float64(delete_results["success"]) / float64(delete_results["total"]))
}

func main() {
    if len(os.Args) < 2 {
        fmt.Printf("Usage: %s <path to log file>\n", os.Args[0])
    } else {
        parse_file(os.Args[1])
    }
}
