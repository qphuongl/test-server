{{- define "sa.GKEAnnotations" }}
{{- if .Values.gke }}
iam.gke.io/gcp-service-account: {{ .Values.serviceAccount.name }}@{{ .Values.gke.projectId }}.iam.gserviceaccount.com
{{- end }}
{{- if .Values.serviceAccount.annotations }}
{{- toYaml .Values.serviceAccount.annotations }}
{{- end }}
{{- end }}