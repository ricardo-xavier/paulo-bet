package xavier.ricardo.paulobet.tasks;

import android.app.ProgressDialog;
import android.os.AsyncTask;

import java.io.IOException;

import okhttp3.Call;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import xavier.ricardo.paulobet.Config;
import xavier.ricardo.paulobet.RankingActivity;

public class RankingTask extends AsyncTask<String, Void, Response> {
    private ProgressDialog progress;
    private RankingActivity context;
    private String league;
    private String user;
    private String token;

    public RankingTask(RankingActivity context, String league, String user, String token) {
        this.context = context;
        this.league = league;
        this.user = user;
        this.token = token;
    }

    @Override
    protected void onPreExecute() {
        progress = new ProgressDialog(context);
        progress.setMessage("Aguarde...");
        progress.show();
        super.onPreExecute();
    }

    @Override
    protected void onPostExecute(Response response) {
        super.onPostExecute(response);
        progress.dismiss();
        try {
            context.onTaskResponse(response);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    @Override
    protected void onProgressUpdate(Void... values) {
        super.onProgressUpdate(values);
    }

    @Override
    protected Response doInBackground(String... params) {
        try {
            Request request = new Request.Builder().url(Config.URL_RANKING + league + "?token=" + token + "&login=" + user).get().build();
            OkHttpClient client = new OkHttpClient.Builder().build();
            Call call = client.newCall(request);
            return call.execute();
        } catch (IOException e) {
            e.printStackTrace();
            return null;
        }
    }
}
